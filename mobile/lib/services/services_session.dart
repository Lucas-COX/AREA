import 'package:http/http.dart' as http;
import 'dart:convert';
import 'dart:async';
import 'package:shared_preferences/shared_preferences.dart';

class User {
  String username;
  DateTime createdAt;
  DateTime updatedAt;

  User(
      {required this.username,
      required this.createdAt,
      required this.updatedAt});

  User.fromJson(Map<String, dynamic> json)
      : username = json['username'],
        createdAt = DateTime.parse(json['created_at']),
        updatedAt = DateTime.parse(json['updated_at']);
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
    try {
      final prefs = await SharedPreferences.getInstance();
      final token = prefs.getString('area_token');
      if (token == null) {
        return Session(null, false);
      }
      final response = await http.get(Uri.parse('http://167.71.52.187:8080/me'),
          headers: {'Authorization': 'Bearer $token'});
      print('Response status: ${response.statusCode}');
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
