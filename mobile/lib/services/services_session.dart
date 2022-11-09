<<<<<<< HEAD
import 'package:http/http.dart' as http;
import 'dart:convert';
import 'dart:async';
import 'package:shared_preferences/shared_preferences.dart';
=======
import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
import 'package:mobil/routes/services/services/services_services.dart';
import 'dart:convert';
import 'dart:async';
import 'package:shared_preferences/shared_preferences.dart';
import '../routes/home/services/service_triggers.dart';

class Trigger {
  num id;
  String title;
  String description;
  String createdAt;
  String updatedAt;
  String reactionData;
  String actionData;
  String action;
  String reaction;
  String actionService;
  String reactionService;
  bool active;

  Trigger(
      {required this.id,
      required this.title,
      required this.description,
      required this.createdAt,
      required this.updatedAt,
      required this.reactionData,
      required this.actionData,
      required this.action,
      required this.reaction,
      required this.actionService,
      required this.reactionService,
      required this.active});

  Trigger.fromJson(Map<String, dynamic> json)
      : id = json['id'],
        title = json['title'],
        description = json['description'],
        createdAt = json['created_at'],
        updatedAt = json['updated_at'],
        reactionData = json['reaction_data'],
        actionData = json['action_data'],
        action = json['action'],
        reaction = json['reaction'],
        actionService = json['action_service'],
        reactionService = json['reaction_service'],
        active = json['active'];

  TriggerBody toTriggerBody() {
    return TriggerBody(
        title: title,
        description: description,
        actionData: actionData,
        reactionData: reactionData,
        action: action,
        reaction: reaction,
        actionService: actionService,
        reactionService: reactionService,
        active: active);
  }
}
>>>>>>> 8dc2ef7 (feat(mobile): creation of a functional flutter client (#72))

class User {
  String username;
  DateTime createdAt;
  DateTime updatedAt;
<<<<<<< HEAD
=======
  List<Trigger> triggers = [];
  List<String> services = [];
>>>>>>> 8dc2ef7 (feat(mobile): creation of a functional flutter client (#72))

  User(
      {required this.username,
      required this.createdAt,
      required this.updatedAt});

  User.fromJson(Map<String, dynamic> json)
      : username = json['username'],
        createdAt = DateTime.parse(json['created_at']),
<<<<<<< HEAD
        updatedAt = DateTime.parse(json['updated_at']);
=======
        updatedAt = DateTime.parse(json['updated_at']),
        triggers = json['triggers'] == null
            ? []
            : (json['triggers'] as List)
                .map((trigger) => Trigger.fromJson(trigger))
                .toList(),
        services = json['services'] == null
            ? []
            : (json['services'] as List)
                .map((service) => service.toString())
                .toList();
>>>>>>> 8dc2ef7 (feat(mobile): creation of a functional flutter client (#72))
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
<<<<<<< HEAD
    try {
      final prefs = await SharedPreferences.getInstance();
=======
    String url = const String.fromEnvironment('API_URL');
    try {
      final prefs = await SharedPreferences.getInstance();

>>>>>>> 8dc2ef7 (feat(mobile): creation of a functional flutter client (#72))
      final token = prefs.getString('area_token');
      if (token == null) {
        return Session(null, false);
      }
<<<<<<< HEAD
      final response = await http.get(Uri.parse('http://167.71.52.187:8080/me'),
          headers: {'Authorization': 'Bearer $token'});
      print('Response status: ${response.statusCode}');
=======
      final response = await http.get(Uri.parse('$url/me'),
          headers: {'Authorization': 'Bearer $token'});

>>>>>>> 8dc2ef7 (feat(mobile): creation of a functional flutter client (#72))
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
