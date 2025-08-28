{
  description = "Zend Home Manager support";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    home-manager.url = "github:nix-community/home-manager";
    home-manager.inputs.nixpkgs.follows = "nixpkgs";
  };

  outputs =
    {
      self,
      nixpkgs,
      home-manager,
    }:
    let
      system = "x86_64-linux";
      pkgs = import nixpkgs { inherit system; };
      hm = home-manager.lib;
    in
    {
      packages.${system}.zend = import ./nix/default.nix { inherit pkgs; };

      devShells.${system}.zend = pkgs.mkShell rec {
        buildInputs = [
          pkgs.nodejs
          pkgs.go
          pkgs.yarn
          pkgs.git
        ];

        ZEND_DIST = "${self.packages.${system}.zend}/share/zend/dist";

        shellHook = ''
          export ZEND_DIST=${ZEND_DIST}
          echo "ZEND_DIST is set to $ZEND_DIST"
        '';
      };

      homeModules = rec {
        default = {
          source = ./nix/zend.nix;
        };

        zend = default;
      };
    };
}
