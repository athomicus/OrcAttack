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
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_TRIGGER\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"orc\"\n"
  "mask: \"ladder\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      x: 6.0\n"
  "      y: -1.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  data: 36.63831\n"
  "  data: 43.46314\n"
  "  data: 10.0\n"
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
    x: -11.0
    y: 21.0
  }
  scale {
    x: 0.6
    y: 0.6
    z: 0.6
  }
}
embedded_components {
  id: "sprite_attack"
  type: "sprite"
  data: "default_animation: \"attack\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "size {\n"
  "  x: 400.0\n"
  "  y: 300.0\n"
  "}\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/images/orc_enemy/Orc.atlas\"\n"
  "}\n"
  ""
  position {
    x: -42.0
    y: 34.0
  }
  rotation {
    z: 0.054643765
    w: 0.9985059
  }
  scale {
    x: 0.73
    y: 0.73
    z: 0.73
  }
}
