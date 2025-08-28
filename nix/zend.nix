{
  config,
  pkgs,
  lib,
  ...
}:

let
  positionKeywords = [
    "top"
    "bottom"
    "left"
    "right"
    "top-left"
    "top-right"
    "bottom-left"
    "bottom-right"
    "center"
  ];

  positionType = lib.types.oneOf [
    (lib.types.enum positionKeywords)
    (lib.types.str)
  ];

  zendPackage = import ./default.nix { inherit pkgs; };
  cfg = config.programs.zend;
in
{
  options.programs.zend.enable = lib.mkOption {
    type = lib.types.bool;
    default = false;
    description = "Enable Zend homepage";
  };

  options.programs.zend.settings = lib.mkOption {
    type = lib.types.attrsOf (
      lib.types.submodule {
        options = {
          dist = lib.mkOption {
            type = lib.types.path;
            default = "./dist";
            description = "Frontend build path";
          };

          random = lib.mkOption {
            type = lib.types.attrsOf (
              lib.types.submodule {
                options = {
                  enabled = lib.mkOption {
                    type = lib.types.bool;
                    default = true;
                  };
                  use = lib.mkOption {
                    type = lib.types.enum [
                      "color"
                      "image"
                    ];
                    default = "color";
                  };
                  interval = lib.mkOption {
                    type = lib.types.int;
                    default = 10;
                  };
                  max = lib.mkOption {
                    type = lib.types.int;
                    default = 100;
                  };
                };
              }
            );
            default = { };
          };

          image = lib.mkOption {
            type = lib.types.attrsOf (
              lib.types.submodule {
                options = {
                  enabled = lib.mkOption {
                    type = lib.type.bool;
                    default = false;
                  };
                  path = lib.mkOption {
                    type = lib.types.path;
                    default = "";
                  };
                  path_list = lib.mkOption {
                    type = lib.types.listOf lib.types.path;
                    default = [ ];
                  };
                };
              }
            );
            default = { };
          };

          resize = lib.mkOption {
            type = lib.types.attrsOf (
              lib.types.submodule {
                options = {
                  mode = lib.mkOption {
                    type = lib.types.enum [
                      "cover"
                      "contain"
                    ];
                    default = "cover";
                  };
                  position = lib.mkOption {
                    type = positionType;
                    default = "center";
                  };
                  repeat = lib.mkOption {
                    type = lib.types.enum [
                      "no-repeat"
                      "repeat"
                      "repeat-x"
                      "repeat-y"
                    ];
                    default = "no-repeat";
                  };
                };
              }
            );
            default = { };
          };

          clock = lib.mkOption {
            type = lib.types.attrsOf (
              lib.types.submodule {
                options = {
                  enabled = lib.mkOption {
                    type = lib.types.bool;
                    default = true;
                  };
                  position = lib.mkOption {
                    type = positionType;
                    default = "top-right";
                  };
                };
              }
            );
            default = { };
          };

          weather = lib.mkOption {
            type = lib.types.attrsOf (
              lib.types.submodule {
                options = {
                  enabled = lib.mkOption {
                    type = lib.types.bool;
                    default = true;
                  };
                  position = lib.mkOption {
                    type = positionType;
                    default = "top-left";
                  };
                  api_key = lib.mkOption {
                    type = lib.types.str;
                    default = "";
                  };
                };
              }
            );
            default = { };
          };

          music = lib.mkOption {
            type = lib.types.attrsOf (
              lib.types.submodule {
                options = {
                  enabled = lib.mkOption {
                    type = lib.types.bool;
                    default = false;
                  };
                  position = lib.mkOption {
                    type = positionType;
                    default = "bottom-left";
                  };
                  local = lib.mkOption {
                    type = lib.types.attrsOf (
                      lib.types.submodule {
                        options = {
                          enabled = lib.mkOption {
                            type = lib.types.bool;
                            default = false;
                          };
                          path = lib.mkOption {
                            type = lib.types.path;
                            default = "";
                          };
                          path_list = lib.mkOption {
                            type = lib.types.listOf lib.types.path;
                            default = [ ];
                          };
                          shuffle = lib.mkOption {
                            type = lib.types.bool;
                            default = true;
                          };
                        };
                      }
                    );
                    default = { };
                  };
                };
              }
            );
            default = { };
          };
        };
      }
    );
    default = { };
    description = "Full Zend configuration with strict type validation for paths and positions";
  };

  config = lib.mkIf cfg.enable {
    # Install Zend package
    home.packages = [ zendPackage ];

    # Generate config.yaml in user's config directory
    home.file."${config.home.homeDirectory}/.config/zend/config.yaml" = {
      text = builtins.toString (
        lib.optionalString (!lib.isNull cfg.config) (
          let
            jsonText = builtins.toJSON cfg.config;
          in
          builtins.concatStringsSep "\n" [
            "# Generated by Home Manager"
            jsonText
          ]
        )
      );
    };

    # Optional systemd user service for Zend daemon
    systemd.user.services.zend = {
      description = "Zend homepage daemon";
      wantedBy = [ "default.target" ];
      serviceConfig.ExecStart = "${zendPackage}/bin/zend";
      restart = "always";
      enabled = true;
    };
  };
}
