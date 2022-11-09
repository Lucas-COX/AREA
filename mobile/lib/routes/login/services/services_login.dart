import 'dart:async';

<<<<<<< HEAD
=======
import 'package:flutter/material.dart';
>>>>>>> 8dc2ef7 (feat(mobile): creation of a functional flutter client (#72))
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';

class ServicesLogin {
  static String url = dotenv.env['API_URL']!;

  static Future login(String username, String password) async {
    var completer = Completer();
<<<<<<< HEAD
    try {
      final response =
          await http.post(Uri.parse('http://167.71.52.187:8080/login'),
              headers: <String, String>{
                'Content-Type': 'application/json; charset=UTF-8',
              },
              body: jsonEncode(<String, String>{
                'username': username,
                'password': password,
              }));
      print('Response status: ${response.statusCode}');
      print('Response body: ${response.body}');
      completer.complete(response);
    } catch (e) {
      print(e);
=======
    String url = const String.fromEnvironment('API_URL');
    try {
      final response = await http.post(Uri.parse('$url/login'),
          headers: <String, String>{
            'Content-Type': 'application/json; charset=UTF-8',
          },
          body: jsonEncode(<String, String>{
            'username': username,
            'password': password,
          }));
      completer.complete(response);
    } catch (e) {
      debugPrint(e.toString());
>>>>>>> 8dc2ef7 (feat(mobile): creation of a functional flutter client (#72))
    }
    return completer.future;
  }
}
