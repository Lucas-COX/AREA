import 'dart:async';

import 'package:flutter/cupertino.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';

import 'package:shared_preferences/shared_preferences.dart';

class TriggerBody {
  String title;
  String description;
  String reactionData;
  String actionData;
  String action;
  String reaction;
  String actionService;
  String reactionService;
  bool active;

  TriggerBody(
      {required this.title,
      required this.description,
      required this.actionData,
      required this.reactionData,
      required this.action,
      required this.reaction,
      required this.actionService,
      required this.reactionService,
      required this.active});

  Map<String, dynamic> toJson() => {
        'title': title,
        'description': description,
        'action_data': actionData,
        'reaction_data': reactionData,
        'action': action,
        'reaction': reaction,
        'action_service': actionService,
        'reaction_service': reactionService,
        'active': active,
      };
}

class TriggersService {
  static String url = const String.fromEnvironment("API_URL");

  static Future post() async {
    var completer = Completer();
    final TriggerBody triggerBody = TriggerBody(
        title: "New trigger",
        description: "",
        actionData: "",
        reactionData: "",
        action: "receive",
        reaction: "receive",
        actionService: "",
        reactionService: "",
        active: true);

    try {
      final prefs = await SharedPreferences.getInstance();
      final token = prefs.getString('area_token');

      final response = await http.post(Uri.parse('$url/triggers'),
          headers: <String, String>{
            'Authorization': 'Bearer $token',
            'Content-Type': 'application/json; charset=UTF-8',
          },
          body: jsonEncode(triggerBody.toJson()));
      completer.complete(response);
    } catch (e) {
      completer.completeError(e);
    }
    return completer.future;
  }

  static Future delete(num id) async {
    var completer = Completer();
    try {
      final prefs = await SharedPreferences.getInstance();
      final token = prefs.getString('area_token');
      final response = await http
          .delete(Uri.parse('$url/triggers/$id'), headers: <String, String>{
        'Authorization': 'Bearer $token',
      });
      completer.complete(response);
    } catch (e) {
      completer.completeError(e);
    }
    return completer.future;
  }

  static Future put(TriggerBody triggerBody, num id) async {
    var completer = Completer();
    try {
      final prefs = await SharedPreferences.getInstance();
      final token = prefs.getString('area_token');
      final response = await http.put(Uri.parse('$url/triggers/$id'),
          headers: <String, String>{
            'Authorization': 'Bearer $token',
            'Content-Type': 'application/json; charset=UTF-8',
          },
          body: jsonEncode(triggerBody.toJson()));
      completer.complete(response);
    } catch (e) {
      completer.completeError(e);
    }
    return completer.future;
  }

  static Future update(TriggerBody trigger, num id) async {
    var completer = Completer();
    try {
      final prefs = await SharedPreferences.getInstance();
      final token = prefs.getString('area_token');
      final response = await http.put(Uri.parse('$url/triggers/$id'),
          headers: <String, String>{
            'Authorization': 'Bearer $token',
            'Content-Type': 'application/json; charset=UTF-8',
          },
          body: jsonEncode(trigger.toJson()));
      completer.complete(response);
    } catch (e) {
      completer.completeError(e);
    }
    return completer.future;
  }
}
