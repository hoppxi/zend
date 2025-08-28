with import <nixpkgs> { };

mkShell {
  buildInputs = [
    go
    nodejs_20
    git
  ];

  ZEND_DIST = "${import ./nix/default.nix { pkgs = pkgs; }}/share/zend/dist";

  shellHook = ''
    export ZEND_DIST=$ZEND_DIST
    echo "ZEND_DIST is set to $ZEND_DIST"
  '';
}
