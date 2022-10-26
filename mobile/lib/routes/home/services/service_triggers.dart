import 'dart:async';

import 'package:flutter/cupertino.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';

import 'package:shared_preferences/shared_preferences.dart';

class TriggerReactionBody {
  String type;
  String action;
  String token;

  TriggerReactionBody(
      {required this.type, required this.action, required this.token});

  Map<String, dynamic> toJson() {
    return {
      'type': type,
      'action': action,
      'token': token,
    };
  }
}

class TriggerActionBody {
  String type;
  String event;
  String token;

  TriggerActionBody(
      {required this.type, required this.event, required this.token});

  Map<String, dynamic> toJson() => {
        'type': type,
        'event': event,
        'token': token,
      };
}

class TriggerBody {
  String title;
  String description;
  num actionId;
  num reactionId;
  String reactionData;
  String actionData;

  TriggerBody(
      {required this.title,
      required this.description,
      required this.actionId,
      required this.reactionId,
      required this.actionData,
      required this.reactionData});

  Map<String, dynamic> toJson() => {
        'title': title,
        'description': description,
        'action_id': actionId,
        'reaction_id': reactionId,
        'action_data': actionData,
        'reaction_data': reactionData,
      };
}

class TriggersService {
  static String url = const String.fromEnvironment("API_URL");

  static Future post() async {
    var completer = Completer();
    final TriggerBody triggerBody = TriggerBody(
        title: "New trigger",
        description: "",
        actionId: 1,
        reactionId: 1,
        actionData: "",
        reactionData: "");
    try {
      final prefs = await SharedPreferences.getInstance();
      final token = prefs.getString('area_token');
      debugPrint(token);

      final response = await http.post(Uri.parse('$url/triggers'),
          headers: <String, String>{
            'Authorization': 'Bearer $token',
            'Content-Type': 'application/json; charset=UTF-8',
          },
          body: jsonEncode(triggerBody.toJson()));
      debugPrint('Response status: ${response.statusCode}');
      debugPrint('Response body: ${response.body}');
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
      debugPrint('Response status: ${response.statusCode}');
      debugPrint('Response body: ${response.body}');
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
      debugPrint('Response status: ${response.statusCode}');
      debugPrint('Response body: ${response.body}');
      completer.complete(response);
    } catch (e) {
      completer.completeError(e);
    }
    return completer.future;
  }
}
