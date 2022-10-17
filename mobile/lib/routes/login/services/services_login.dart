import 'dart:async';

import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';

class ServicesLogin {
  static String url = dotenv.env['API_URL']!;

  static Future login(String username, String password) async {
    var completer = Completer();
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
      print('Response status: ${response.statusCode}');
      print('Response body: ${response.body}');
      completer.complete(response);
    } catch (e) {
      print(e);
    }
    return completer.future;
  }
}
