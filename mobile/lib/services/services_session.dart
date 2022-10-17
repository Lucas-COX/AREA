import 'package:http/http.dart' as http;
import 'dart:convert';
import 'dart:async';
import 'package:shared_preferences/shared_preferences.dart';
import '../routes/home/services/service_triggers.dart';

class TriggerAction {
  num id;
  String type;
  String event;
  num triggerId;

  TriggerAction(this.id, this.type, this.event, this.triggerId);
  TriggerAction.fromJson(Map<String, dynamic> json)
      : id = json['id'],
        type = json['type'],
        event = json['event'],
        triggerId = json['trigger_id'];
}

class TriggerReaction {
  num id;
  String type;
  String action;
  num triggerId;
  String token;

  TriggerReaction(this.id, this.type, this.action, this.triggerId, this.token);
  TriggerReaction.fromJson(Map<String, dynamic> json)
      : id = json['id'],
        type = json['type'],
        action = json['action'],
        triggerId = json['trigger_id'],
        token = json['token'];
}

class Trigger {
  num id;
  String title;
  String description;
  TriggerAction action;
  TriggerReaction reaction;
  String createdAt;
  String updatedAt;

  Trigger(
      {required this.id,
      required this.title,
      required this.description,
      required this.action,
      required this.reaction,
      required this.createdAt,
      required this.updatedAt});

  Trigger.fromJson(Map<String, dynamic> json)
      : id = json['id'],
        title = json['title'],
        description = json['description'],
        action = TriggerAction.fromJson(json['action']),
        reaction = TriggerReaction.fromJson(json['reaction']),
        createdAt = json['created_at'],
        updatedAt = json['updated_at'];

  TriggerBody toTriggerBody() {
    return TriggerBody(
        title: title,
        description: description,
        action: TriggerActionBody(
            type: action.type, event: action.event, token: ''),
        reaction: TriggerReactionBody(
            type: reaction.type,
            action: reaction.action,
            token: reaction.token));
  }
}

class User {
  String username;
  DateTime createdAt;
  DateTime updatedAt;
  List<Trigger> triggers = [];

  User(
      {required this.username,
      required this.createdAt,
      required this.updatedAt});

  User.fromJson(Map<String, dynamic> json)
      : username = json['username'],
        createdAt = DateTime.parse(json['created_at']),
        updatedAt = DateTime.parse(json['updated_at']),
        triggers = json['triggers'] == null
            ? []
            : (json['triggers'] as List)
                .map((trigger) => Trigger.fromJson(trigger))
                .toList();
}

class Session {
  User? user;
  bool isLoggedIn = false;

  Session(this.user, this.isLoggedIn);
  Session.fromJson(Map<String, dynamic> json)
      : user = User.fromJson(json),
        isLoggedIn = true;
}

class ServicesSession {
  static Future<Session> get() async {
    var completer = Completer<Session>();
    String url = const String.fromEnvironment('API_URL');
    try {
      final prefs = await SharedPreferences.getInstance();
      final token = prefs.getString('area_token');
      if (token == null) {
        return Session(null, false);
      }
      final response = await http.get(Uri.parse('$url/me'),
          headers: {'Authorization': 'Bearer $token'});
      print('Response status session: ${response.statusCode}');
      final json = jsonDecode(response.body);
      if (!json.containsKey('me')) {
        completer.complete(Session(null, false));
      } else {
        completer.complete(Session.fromJson(json['me']));
      }
    } catch (e) {
      completer.completeError(e);
    }
    return completer.future;
  }
}
