import 'package:flutter/material.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:mobil/routes/login/login.dart';
import 'package:shared_preferences/shared_preferences.dart';
import 'dart:async';
import 'dart:convert';
import 'package:http/http.dart' as http;

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
        createdAt = DateTime.parse(json['createdAt']),
        updatedAt = DateTime.parse(json['updatedAt']);
}

class Session {
  bool authenticated;
  User? user;

  Session({required this.authenticated, required this.user});
  Session.load()
      : authenticated = false,
        user = null {
    dotenv.load();
    final uri = dotenv.env['API_URL'] ?? "";
    http.get(Uri.parse('http://$uri/me'), headers: <String, String>{
      'Content-Type': 'application/json; charset=UTF-8',
    }).then((response) {
      if (response.statusCode == 200) {
        authenticated = true;
        user = User.fromJson(jsonDecode(response.body));
      }
    });
  }

  bool _isAuthenticated() {
    return authenticated;
  }

  void _setAuthenticated(bool value) {
    authenticated = value;
  }

  User? _getUser() {
    return user;
  }

  void _setUser(User? value) {
    user = value;
  }
}

class HomePage extends StatefulWidget {
  const HomePage({Key? key, required this.title}) : super(key: key);

  final String title;

  @override
  State<HomePage> createState() => _HomePageState();
}

class _HomePageState extends State<HomePage> {
  final Session _session = Session.load();

  User? _getUser() {
    return _session._getUser();
  }

  Session _getSession() {
    return _session;
  }

  @override
  Widget build(BuildContext context) {
    if (_getSession()._isAuthenticated()) {
      Navigator.pushNamed(context, "/login");
    }
    return Scaffold(
      appBar: AppBar(
        title: Text(widget.title),
      ),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: <Widget>[
            Text(
              'Welcome ${_getUser()?.username}',
            ),
          ],
        ),
      ),
    );
  }
}
