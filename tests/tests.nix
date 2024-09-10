{ ... }:
{
  perSystem = { pkgs, self', system, ... }: {
    packages = {
      tests = pkgs.go.stdenv.mkDerivation {
        name = "github.com/cosmos/cosmos-sdk/tests";
        buildInputs = [ pkgs.go pkgs.gotestsum pkgs.git pkgs.which ];
        src = ./.;
        doCheck = true;
        dontInstall = true;
        dontBuild = true;
        checkPhase = ''
          export HOME=$(pwd)
          export GOFLAGS="-mod=vendor -tags=\"cgo,ledger,test_ledger_mock,norace\""
          export GOPRIVATE="github.com/unionlabs/*";
          export GOWORK="off"

          echo "Running test for github.com/cosmos/cosmos-sdk/tests sub module..."
          go version

          gotestsum  ./...

          echo "Finished running sub module tests."
          
          touch $out
        '';
      };
    };
  };
}
