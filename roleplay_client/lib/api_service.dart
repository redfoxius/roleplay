import 'dart:convert';
import 'package:http/http.dart' as http;

class ApiService {
  static const String baseUrl = 'http://localhost:8080/api'; // Update as needed

  // Create a new character
  Future<void> createCharacter(String name, String charClass) async {
    final response = await http.post(
      Uri.parse('$baseUrl/character/create'),
      headers: {'Content-Type': 'application/json'},
      body: jsonEncode({'name': name, 'class': charClass}),
    );

    if (response.statusCode != 200) {
      throw Exception('Failed to create character: ${response.body}');
    }
  }

  // Get the world map
  Future<WorldMap> getWorldMap() async {
    final response = await http.get(Uri.parse('$baseUrl/world/map'));

    if (response.statusCode != 200) {
      throw Exception('Failed to get world map: ${response.body}');
    }

    return WorldMap.fromJson(jsonDecode(response.body));
  }

  // Start a battle
  Future<Battle> startBattle({required String type, required List<String> participants, Map<String, List<String>>? teams, String? terrain, String? weather}) async {
    final response = await http.post(
      Uri.parse('$baseUrl/combat/start'),
      headers: {'Content-Type': 'application/json'},
      body: jsonEncode({
        'type': type,
        'participants': participants,
        'teams': teams,
        'terrain': terrain,
        'weather': weather,
      }),
    );

    if (response.statusCode != 200) {
      throw Exception('Failed to start battle: ${response.body}');
    }

    return Battle.fromJson(jsonDecode(response.body));
  }

  // Add more methods as needed for other endpoints
} 