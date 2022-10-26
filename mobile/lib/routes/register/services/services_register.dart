import 'package:flutter/cupertino.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';
import 'dart:async';

class ServicesRegister {
  static Future register(String username, String password) async {
    var completer = Completer();
    String url = const String.fromEnvironment('API_URL');
    debugPrint(url);
    try {
      final response = await http.post(Uri.parse('$url/register'),
          headers: <String, String>{
            'Content-Type': 'application/json; charset=UTF-8',
          },
          body: jsonEncode(<String, String>{
            'username': username,
            'password': password,
          }));
      debugPrint('Response status: ${response.statusCode}');
      debugPrint('Response body: ${response.body}');
      completer.complete(response);
    } catch (e) {
      debugPrint(e.toString());
    }
    return completer.future;
  }
}
