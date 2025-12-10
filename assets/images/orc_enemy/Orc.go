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
}
embedded_components {
  id: "sprite_fade"
  type: "sprite"
  data: "default_animation: \"walkToClimb\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/images/orc_enemy/Orc.atlas\"\n"
  "}\n"
  ""
  position {
    x: -27.0
    y: 69.0
  }
  scale {
    x: 1.7
    y: 1.7
    z: 1.7
  }
}
