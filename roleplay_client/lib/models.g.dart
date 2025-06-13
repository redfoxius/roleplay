// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'models.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

Character _$CharacterFromJson(Map<String, dynamic> json) => Character(
  id: json['id'] as String,
  name: json['name'] as String,
  charClass: json['charClass'] as String?,
  level: (json['level'] as num?)?.toInt(),
  health: (json['health'] as num?)?.toInt(),
  steamPower: (json['steamPower'] as num?)?.toInt(),
);

Map<String, dynamic> _$CharacterToJson(Character instance) => <String, dynamic>{
  'id': instance.id,
  'name': instance.name,
  'charClass': instance.charClass,
  'level': instance.level,
  'health': instance.health,
  'steamPower': instance.steamPower,
};

WorldMap _$WorldMapFromJson(Map<String, dynamic> json) => WorldMap(
  width: (json['width'] as num).toInt(),
  height: (json['height'] as num).toInt(),
  locations: (json['locations'] as List<dynamic>)
      .map((e) => Location.fromJson(e as Map<String, dynamic>))
      .toList(),
);

Map<String, dynamic> _$WorldMapToJson(WorldMap instance) => <String, dynamic>{
  'width': instance.width,
  'height': instance.height,
  'locations': instance.locations,
};

Location _$LocationFromJson(Map<String, dynamic> json) => Location(
  id: json['id'] as String,
  name: json['name'] as String,
  type: json['type'] as String,
  description: json['description'] as String,
  terrain: json['terrain'] as String,
  region: json['region'] as String,
  coordinates: Coordinates.fromJson(
    json['coordinates'] as Map<String, dynamic>,
  ),
);

Map<String, dynamic> _$LocationToJson(Location instance) => <String, dynamic>{
  'id': instance.id,
  'name': instance.name,
  'type': instance.type,
  'description': instance.description,
  'terrain': instance.terrain,
  'region': instance.region,
  'coordinates': instance.coordinates,
};

Coordinates _$CoordinatesFromJson(Map<String, dynamic> json) =>
    Coordinates(x: (json['x'] as num).toInt(), y: (json['y'] as num).toInt());

Map<String, dynamic> _$CoordinatesToJson(Coordinates instance) =>
    <String, dynamic>{'x': instance.x, 'y': instance.y};

Battle _$BattleFromJson(Map<String, dynamic> json) => Battle(
  id: json['id'] as String,
  type: json['type'] as String,
  state: json['state'] as String,
);

Map<String, dynamic> _$BattleToJson(Battle instance) => <String, dynamic>{
  'id': instance.id,
  'type': instance.type,
  'state': instance.state,
};

Ability _$AbilityFromJson(Map<String, dynamic> json) => Ability(
  name: json['name'] as String,
  description: json['description'] as String,
  type: json['type'] as String,
  damage: (json['damage'] as num?)?.toInt(),
  healing: (json['healing'] as num?)?.toInt(),
  steamCost: (json['steamCost'] as num?)?.toInt(),
  range: (json['range'] as num?)?.toInt(),
  area: (json['area'] as num?)?.toInt(),
  cooldown: (json['cooldown'] as num?)?.toInt(),
);

Map<String, dynamic> _$AbilityToJson(Ability instance) => <String, dynamic>{
  'name': instance.name,
  'description': instance.description,
  'type': instance.type,
  'damage': instance.damage,
  'healing': instance.healing,
  'steamCost': instance.steamCost,
  'range': instance.range,
  'area': instance.area,
  'cooldown': instance.cooldown,
};
