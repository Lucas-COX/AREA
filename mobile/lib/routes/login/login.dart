import 'dart:convert';

import 'package:shared_preferences/shared_preferences.dart';
import 'services/services_login.dart';
import 'package:flutter/material.dart';

class AccueilPage extends StatefulWidget {
  const AccueilPage({Key? key, required this.title}) : super(key: key);

  final String title;

  @override
  State<AccueilPage> createState() => _AccueilPageState();
}

class _AccueilPageState extends State<AccueilPage> {
  String username = '', password = '';

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: const Color.fromRGBO(255, 255, 255, 1),
      appBar: AppBar(
        centerTitle: true,
        backgroundColor: const Color.fromRGBO(255, 255, 255, 1),
        title: const Text('Area',
            style: TextStyle(color: Color.fromRGBO(255, 255, 255, 1))),
        elevation: 0,
        leading: const Text(''),
      ),
      body: Form(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            const Text(
              'Welcome to Area',
              style: TextStyle(
                fontSize: 30,
                fontWeight: FontWeight.bold,
                color: Color.fromRGBO(206, 13, 13, 1),
              ),
            ),
            const SizedBox(height: 80),
            Padding(
              padding: const EdgeInsets.symmetric(horizontal: 20),
              child: TextFormField(
                decoration: const InputDecoration(
                  contentPadding: EdgeInsets.all(14),
                  labelText: 'Username',
                  focusedBorder: OutlineInputBorder(
                    borderSide: BorderSide(
                        color: Color.fromRGBO(12, 169, 12, 1), width: 2.0),
                    borderRadius: BorderRadius.all(Radius.circular(10)),
                  ),
                  border: OutlineInputBorder(
                    borderRadius: BorderRadius.all(Radius.circular(10)),
                  ),
                ),
                onChanged: (value) => setState(() {
                  username = value;
                }),
              ),
            ),
            const SizedBox(height: 20),
            Padding(
              padding: const EdgeInsets.symmetric(horizontal: 20),
              child: TextFormField(
                obscureText: true,
                decoration: const InputDecoration(
                  hoverColor: Color.fromRGBO(12, 169, 12, 1),
                  labelText: 'Password',
                  contentPadding: EdgeInsets.all(14),
                  focusedBorder: OutlineInputBorder(
                    borderSide: BorderSide(
                        color: Color.fromRGBO(12, 169, 12, 1), width: 2.0),
                    borderRadius: BorderRadius.all(Radius.circular(10)),
                  ),
                  border: OutlineInputBorder(
                    borderRadius: BorderRadius.all(Radius.circular(10)),
                    borderSide:
                        BorderSide(color: Color.fromRGBO(12, 169, 12, 1)),
                  ),
                ),
                onChanged: (value) => setState(() {
                  password = value;
                }),
              ),
            ),
            const SizedBox(height: 20),
            ElevatedButton(
                style: ElevatedButton.styleFrom(
                  backgroundColor: const Color.fromRGBO(12, 169, 12, 1),
                  padding:
                      const EdgeInsets.symmetric(horizontal: 45, vertical: 10),
                ),
                onPressed: () async {
                  final response = await ServicesLogin.login(
                      username.trim(), password.trim());
                  if (response.statusCode == 200) {
                    final prefs = await SharedPreferences.getInstance();
                    prefs.setString(
                        'area_token', jsonDecode(response.body)['token']);
                    if (prefs.getString('area_token') != null) {
                      Navigator.pushNamed(context, '/');
                    }
                  } else {
                    showDialog(
                        context: context,
                        builder: (BuildContext context) {
                          return AlertDialog(
                            title: const Text('Error'),
                            content: const Text('Wrong Password or Username'),
                            actions: [
                              TextButton(
                                onPressed: () {
                                  Navigator.of(context).pop();
                                },
                                child: const Text('OK'),
                              ),
                            ],
                          );
                        });
                  }
                },
                child: const Text(
                  'Login',
                  style: TextStyle(
                    fontSize: 15,
                    fontWeight: FontWeight.bold,
                    color: Color.fromRGBO(64, 61, 57, 1),
                  ),
                )),
            TextButton(
                onPressed: () {
                  Navigator.pushNamed(context, '/register');
                },
                child: const Text(
                  'Sign in',
                  style: TextStyle(color: Color.fromRGBO(64, 61, 57, 1)),
                )),
          ],
        ),
      ),
    );
  }
}
