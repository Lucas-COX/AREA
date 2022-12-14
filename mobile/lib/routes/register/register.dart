import 'dart:convert';
import 'package:flutter/material.dart';
import 'package:shared_preferences/shared_preferences.dart';
import 'services/services_register.dart';

class RegisterPage extends StatelessWidget {
  const RegisterPage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    String username = '';
    String password = '';
    return Scaffold(
      backgroundColor: const Color.fromRGBO(255, 255, 255, 1),
      appBar: AppBar(
        centerTitle: true,
        title: const Text('Register'),
        leading: const Text(''),
        backgroundColor: const Color.fromRGBO(255, 255, 255, 1),
      ),
      body: Form(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            const SizedBox(height: 30),
            const Text(
              'Welcome to Area',
              style: TextStyle(
                  color: Color.fromRGBO(206, 13, 13, 1),
                  fontSize: 25,
                  fontWeight: FontWeight.bold),
            ),
            const SizedBox(height: 17),
            const Text(
              'Please enter your credentials',
              style: TextStyle(
                  color: Color.fromRGBO(206, 13, 13, 1),
                  fontSize: 15,
                  fontWeight: FontWeight.bold),
            ),
            const SizedBox(height: 20),
            const Text(
              'Username',
              style: TextStyle(
                  color: Color.fromRGBO(206, 13, 13, 1),
                  fontSize: 15,
                  fontWeight: FontWeight.bold),
            ),
            const SizedBox(height: 17),
            Padding(
              padding: const EdgeInsets.symmetric(horizontal: 20),
              child: TextFormField(
                decoration: const InputDecoration(
                  hoverColor: Color.fromRGBO(206, 13, 13, 1),
                  contentPadding: EdgeInsets.all(14),
                  focusedBorder: OutlineInputBorder(
                    borderSide: BorderSide(
                        color: Color.fromRGBO(12, 169, 12, 1), width: 2.0),
                    borderRadius: BorderRadius.all(Radius.circular(10)),
                  ),
                  border: OutlineInputBorder(
                    borderRadius: BorderRadius.all(Radius.circular(10)),
                  ),
                  labelText: 'Username',
                  icon: Icon(Icons.person),
                  suffix: Icon(Icons.check),
                ),
                onChanged: (value) {
                  username = value;
                },
                style: const TextStyle(
                    color: Color.fromRGBO(206, 13, 13, 1),
                    fontSize: 15,
                    fontWeight: FontWeight.bold),
              ),
            ),
            const SizedBox(height: 17),
            const Text(
              'Password',
              style: TextStyle(
                  color: Color.fromRGBO(206, 13, 13, 1),
                  fontSize: 15,
                  fontWeight: FontWeight.bold),
            ),
            const SizedBox(height: 17),
            Padding(
              padding: const EdgeInsets.symmetric(horizontal: 20),
              child: TextFormField(
                obscureText: true,
                decoration: const InputDecoration(
                  hoverColor: Color.fromRGBO(206, 13, 13, 1),
                  contentPadding: EdgeInsets.all(14),
                  focusedBorder: OutlineInputBorder(
                    borderSide: BorderSide(
                        color: Color.fromRGBO(12, 169, 12, 1), width: 2.0),
                    borderRadius: BorderRadius.all(Radius.circular(10)),
                  ),
                  border: OutlineInputBorder(
                    borderRadius: BorderRadius.all(Radius.circular(10)),
                    borderSide:
                        BorderSide(color: Color.fromRGBO(206, 13, 13, 1)),
                  ),
                  labelText: 'Password',
                  icon: Icon(Icons.lock),
                  suffix: Icon(Icons.check),
                ),
                onChanged: (value) {
                  password = value;
                },
                style: const TextStyle(
                    color: Color.fromRGBO(206, 13, 13, 1),
                    fontSize: 15,
                    fontWeight: FontWeight.bold),
              ),
            ),
            const SizedBox(height: 25),
            ElevatedButton(
              onPressed: () async {
                final response = await ServicesRegister.register(
                    username.trim(), password.trim());
                print(response.statusCode);
                if (response.statusCode == 200) {
                  final prefs = await SharedPreferences.getInstance();
                  prefs.setString(
                      'area_token', jsonDecode(response.body)['token']);
                  if (prefs.getString('area_token') != null) {}
                  Navigator.pushNamed(context, '/');
                } else {
                  ScaffoldMessenger.of(context).showSnackBar(
                    const SnackBar(
                      content: Text('Error'),
                    ),
                  );
                }
              },
              style: ElevatedButton.styleFrom(
                backgroundColor: const Color.fromRGBO(12, 169, 12, 1),
                padding:
                    const EdgeInsets.symmetric(horizontal: 30, vertical: 15),
              ),
              child: const Text('Submit',
                  style: TextStyle(
                      color: Color.fromRGBO(206, 13, 13, 1),
                      fontSize: 18,
                      fontWeight: FontWeight.bold)),
            ),
          ],
        ),
      ),
    );
  }
}
