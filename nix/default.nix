{ pkgs }:

let
  goApp = pkgs.buildGoModule {
    pname = "zend-go";
    version = "1.0.0";

    src = pkgs.lib.cleanSource ./.;

    vendorSha256 = "sha-256ejibibwiejpwepwejwercpiecbwebcjbwebcebicbwehbc";

    subPackages = [ "cmd" ];

    installPhase = ''
      mkdir -p $out/bin
      cp cmd/zend $out/bin/
    '';
  };

  frontend = pkgs.buildNpmPackage {
    pname = "zend-web";
    version = "1.0.0";

    src = ./web;

    packageJSON = ./web/package.json;
    packageLockJSON = ./web/package-lock.json;

    buildPhase = ''
      npm install
      npm run build
      mkdir -p $out/dist
      cp -r dist/* $out/dist/
    '';

    installPhase = ''
      mkdir -p $out/share/zend
      cp -r dist $out/share/zend/
    '';
  };
in
pkgs.stdenv.mkDerivation {
  pname = "zend";
  version = "1.0.0";

  buildInputs = [
    goApp
    frontend
  ];

  phases = [ "installPhase" ];

  installPhase = ''
    mkdir -p $out/bin
    cp -r ${goApp}/bin/* $out/bin/
    cp -r ${frontend}/share/zend/dist $out/share/zend/
  '';

  # Expose ZEND_DIST for shell usage
  shellHook = ''
    export ZEND_DIST=$out/share/zend/dist
    echo "ZEND_DIST set to $ZEND_DIST"
  '';
}
