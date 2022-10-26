import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';
import 'dart:async';
import 'package:shared_preferences/shared_preferences.dart';
import '../routes/home/services/service_triggers.dart';

class Trigger {
  num id;
  String title;
  String description;
  num actionId;
  num reactionId;
  String createdAt;
  String updatedAt;
  String reactionData;
  String actionData;

  Trigger(
      {required this.id,
      required this.title,
      required this.description,
      required this.actionId,
      required this.reactionId,
      required this.createdAt,
      required this.updatedAt,
      required this.reactionData,
      required this.actionData});

  Trigger.fromJson(Map<String, dynamic> json)
      : id = json['id'],
        title = json['title'],
        description = json['description'],
        actionId = json['action_id'],
        reactionId = json['reaction_id'],
        createdAt = json['created_at'],
        updatedAt = json['updated_at'],
        reactionData = json['reaction_data'],
        actionData = json['action_data'];

  TriggerBody toTriggerBody() {
    return TriggerBody(
        title: title,
        description: description,
        actionId: actionId,
        reactionId: reactionId,
        actionData: actionData,
        reactionData: reactionData);
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
    debugPrint(url);
    try {
      final prefs = await SharedPreferences.getInstance();

      final token = prefs.getString('area_token');
      debugPrint('token = $token');
      if (token == null) {
        return Session(null, false);
      }
      final response = await http.get(Uri.parse('$url/me'),
          headers: {'Authorization': 'Bearer $token'});
      debugPrint('Response status session: ${response.statusCode}');
      debugPrint('Response body session: ${response.body}');
      final json = jsonDecode(response.body);
      if (!json.containsKey('me')) {
        completer.complete(Session(null, false));
      } else {
        completer.complete(Session.fromJson(json['me']));
      }
    } catch (e) {
      completer.completeError(e);
    }
    print('completer = ${completer.future}');
    return completer.future;
  }
}
