import 'package:json_annotation/json_annotation.dart';

part 'models.g.dart';

@JsonSerializable()
class Character {
  final String id;
  final String name;
  final String? charClass;
  final int? level;
  final int? health;
  final int? steamPower;
  // TODO: Add attributes, money, inventory, etc.

  Character({
    required this.id,
    required this.name,
    this.charClass,
    this.level,
    this.health,
    this.steamPower,
  });

  factory Character.fromJson(Map<String, dynamic> json) => _$CharacterFromJson(json);
  Map<String, dynamic> toJson() => _$CharacterToJson(this);
}

@JsonSerializable()
class WorldMap {
  final int width;
  final int height;
  final List<Location> locations;
  // TODO: Add regions, etc.

  WorldMap({
    required this.width,
    required this.height,
    required this.locations,
  });

  factory WorldMap.fromJson(Map<String, dynamic> json) => _$WorldMapFromJson(json);
  Map<String, dynamic> toJson() => _$WorldMapToJson(this);
}

@JsonSerializable()
class Location {
  final String id;
  final String name;
  final String type;
  final String description;
  final String terrain;
  final String region;
  final Coordinates coordinates;
  // TODO: Add properties, resources, etc.

  Location({
    required this.id,
    required this.name,
    required this.type,
    required this.description,
    required this.terrain,
    required this.region,
    required this.coordinates,
  });

  factory Location.fromJson(Map<String, dynamic> json) => _$LocationFromJson(json);
  Map<String, dynamic> toJson() => _$LocationToJson(this);
}

@JsonSerializable()
class Coordinates {
  final int x;
  final int y;

  Coordinates({required this.x, required this.y});

  factory Coordinates.fromJson(Map<String, dynamic> json) => _$CoordinatesFromJson(json);
  Map<String, dynamic> toJson() => _$CoordinatesToJson(this);
}

@JsonSerializable()
class Battle {
  final String id;
  final String type;
  final String state;
  // TODO: Add participants, turn order, terrain, weather, etc.

  Battle({
    required this.id,
    required this.type,
    required this.state,
  });

  factory Battle.fromJson(Map<String, dynamic> json) => _$BattleFromJson(json);
  Map<String, dynamic> toJson() => _$BattleToJson(this);
}

@JsonSerializable()
class Ability {
  final String name;
  final String description;
  final String type;
  final int? damage;
  final int? healing;
  final int? steamCost;
  final int? range;
  final int? area;
  final int? cooldown;

  Ability({
    required this.name,
    required this.description,
    required this.type,
    this.damage,
    this.healing,
    this.steamCost,
    this.range,
    this.area,
    this.cooldown,
  });

  factory Ability.fromJson(Map<String, dynamic> json) => _$AbilityFromJson(json);
  Map<String, dynamic> toJson() => _$AbilityToJson(this);
} 