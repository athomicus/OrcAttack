components {
  id: "Orc"
  component: "/main/Orc.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"walking\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/images/orc_enemy/Orc.atlas\"\n"
  "}\n"
  ""
  scale {
    x: 0.4
    y: 0.4
    z: 0.4
  }
}
