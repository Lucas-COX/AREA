import 'dart:async';
import 'dart:convert';
import 'package:shared_preferences/shared_preferences.dart';
import 'package:http/http.dart' as http;
import 'package:url_launcher/url_launcher.dart';

class Action {
  String name;
  String description;

  Action({required this.name, required this.description});

  Action.fromJson(Map<String, dynamic> json)
      : name = json['name'],
        description = json['description'];

  Map<String, dynamic> toJson() => {
        'name': name,
        'description': description,
      };
}

class Reaction {
  String name;
  String description;

  Reaction({required this.name, required this.description});

  Reaction.fromJson(Map<String, dynamic> json)
      : name = json['name'],
        description = json['description'];

  Map<String, dynamic> toJson() => {
        'name': name,
        'description': description,
      };
}

class Service {
  List<Action> actions;
  List<Reaction> reactions;
  String name;

  Service({required this.actions, required this.reactions, required this.name});

  Service.fromJson(Map<String, dynamic> json)
      : actions =
            (json['actions'] as List).map((i) => Action.fromJson(i)).toList(),
        reactions = (json['reactions'] as List)
            .map((i) => Reaction.fromJson(i))
            .toList(),
        name = json['name'];

  Map<String, dynamic> toJson() => {
        'actions': actions,
        'reactions': reactions,
        'name': name,
      };
}

class Services {
  static Future<List<Service>> getServices() async {
    var completer = Completer<List<Service>>();
    try {
      final prefs = await SharedPreferences.getInstance();
      final token = prefs.getString('area_token');
      String url = const String.fromEnvironment('API_URL');
      final response =
          await http.get(Uri.parse('$url/services'), headers: <String, String>{
        'Authorization': 'Bearer $token',
      });
      final json = jsonDecode(response.body);
      if (!json.containsKey('services')) {
        completer.complete([]);
      } else {
        completer.complete((json['services'] as List)
            .map((service) => Service.fromJson(service))
            .toList());
      }
    } catch (e) {
      completer.completeError(e);
    }
    print('completer = ${completer.future}');
    return completer.future;
  }

  static Future getUrl(String name) async {
    var completer = Completer();
    final prefs = await SharedPreferences.getInstance();
    final token = prefs.getString('area_token');
    Codec<String, String> stringToBase64Url = utf8.fuse(base64Url);
    String url =
        '${const String.fromEnvironment('API_URL')}/providers/$name/auth?callback=${stringToBase64Url.encode('https://google.com')}';
    if (token != null) {
      try {
        final response =
            await http.get(Uri.parse(url), headers: <String, String>{
          'Authorization': 'Bearer $token',
        });
        completer.complete(jsonDecode(response.body));
        print('body = ${response.body}');
      } catch (e) {
        print(e.toString());
      }
      return completer.future;
    }
  }

  static Future connexion(String url) async {
    if (!await launchUrl(Uri.parse(url),
        mode: LaunchMode.externalApplication)) {
      throw 'Could not launch $url';
    }
  }
}
