import 'dart:async';
import 'dart:convert';
import 'package:shared_preferences/shared_preferences.dart';
import 'package:http/http.dart' as http;
import 'package:url_launcher/url_launcher.dart';

class ServiceAction {
  String name;
  String description;

  ServiceAction({required this.name, required this.description});

  ServiceAction.fromJson(Map<String, dynamic> json)
      : name = json['name'],
        description = json['description'];

  Map<String, dynamic> toJson() => {
        'name': name,
        'description': description,
      };
}

class ServiceReaction {
  String name;
  String description;

  ServiceReaction({required this.name, required this.description});

  ServiceReaction.fromJson(Map<String, dynamic> json)
      : name = json['name'],
        description = json['description'];

  Map<String, dynamic> toJson() => {
        'name': name,
        'description': description,
      };
}

class Service {
  List<ServiceAction> actions;
  List<ServiceReaction> reactions;
  String name;

  Service({required this.actions, required this.reactions, required this.name});

  Service.fromJson(Map<String, dynamic> json)
      : actions = (json['actions'] as List)
            .map((i) => ServiceAction.fromJson(i))
            .toList(),
        reactions = (json['reactions'] as List)
            .map((i) => ServiceReaction.fromJson(i))
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
    return completer.future;
  }

  static Future getUrl(String name) async {
    var completer = Completer();
    final prefs = await SharedPreferences.getInstance();
    final token = prefs.getString('area_token');
    Codec<String, String> stringToBase64Url = utf8.fuse(base64Url);
    String url =
        '${const String.fromEnvironment('API_URL')}/providers/$name/auth?callback=${stringToBase64Url.encode('${const String.fromEnvironment('API_URL')}/login/done')}';
    url = url.substring(0, url.length - 2);
    if (token != null) {
      try {
        final response =
            await http.get(Uri.parse(url), headers: <String, String>{
          'Authorization': 'Bearer $token',
        });
        completer.complete(jsonDecode(response.body));
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
