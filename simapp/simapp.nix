{ ... }: {
  perSystem = { pkgs, lib, self', system, rev, ... }: {
    packages = {
      simapp = 
      let
        pname = "simd";
        version = "v0.0.1";
        tags = [ "ledger" "netgo" "rocksdb" "grocksdb_no_link" ];
        ldflags = [
          "-X github.com/cosmos/cosmos-sdk/version.Name=${pname}"
          "-X github.com/cosmos/cosmos-sdk/version.AppName=${pname}"
          "-X github.com/cosmos/cosmos-sdk/version.Version=${version}"
          "-X github.com/cosmos/cosmos-sdk/version.BuildTags=${lib.concatStringsSep "," tags}"
          "-X github.com/cosmos/cosmos-sdk/version.Commit=${rev}"
        ];
        static = pkgs.stdenv.hostPlatform.isStatic;
        cgo_ldflags =
          if static then "-lrocksdb -pthread -lstdc++ -ldl -lzstd -lsnappy -llz4 -lbz2 -lz"
          else if pkgs.stdenv.hostPlatform.isWindows then "-lrocksdb-shared"
          else "-lrocksdb -pthread -lstdc++ -ldl";
      in
      pkgs.buildGo123Module {
        inherit ldflags;
        name = pname;
        src = ./.;
        vendorHash = null;
        doCheck = false;
        subPackages = [ "simd" ];
        # checkPhase = ''
        #   export HOME=$(pwd)
        #   export CGO_LDFLAGS="${cgo_ldflags}"
        #   # export GOFLAGS="-tags=\"cgo,ledger,test_ledger_mock,norace\""
        #   export GOWORK="off"

        #   echo "Running test for ${pname}"
        #   go version

        #   go test ./...

        #   echo "Finished running ${pname} tests."
        # '';
        buildInputs = [ pkgs.rocksdb_8_11 ];
        overrideModAttrs = (_: {
          CGO_LDFLAGS = cgo_ldflags;
          GOPRIVATE = "github.com/unionlabs/*";
          GOWORK = "off";
        });
      };
    };
  };
}
