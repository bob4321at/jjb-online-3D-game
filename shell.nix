{ pkgs ? import <nixpkgs> {} }:
pkgs.mkShell {
  buildInputs = with pkgs; [
    go
    gcc
    xorg.libXrandr
    xorg.libXinerama
    xorg.libXcursor
    xorg.libXi
    xorg.libXxf86vm
    wayland-utils
    wayland.dev
    libxkbcommon
    libGL
  ];

  shellHook = ''
     export LD_LIBRARY_PATH=${pkgs.mesa}/lib:${pkgs.libglvnd}/lib:$LD_LIBRARY_PATH
    export LD_LIBRARY_PATH="/run/opengl-driver/lib:/run/opengl-driver-32/lib"
  '';


    CGO_ENABLED = "1";
}

