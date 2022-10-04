import 'package:flutter/material.dart';
import 'dart:async';

import '../../services/services_session.dart';

class HomePage extends StatefulWidget {
  const HomePage({Key? key, required this.title}) : super(key: key);

  final String title;

  @override
  State<HomePage> createState() => _HomePageState();
}

class _HomePageState extends State<HomePage> {
  @override
  Widget build(BuildContext context) {
    return FutureBuilder<Session>(
        future: ServicesSession.get(),
        builder: (BuildContext context, AsyncSnapshot<Session> snapshot) {
          if (!snapshot.hasData) {
            return (const Scaffold(
                body: Center(child: CircularProgressIndicator())));
          }
          final session = snapshot.data;
          if (session?.isLoggedIn != null && !(session?.isLoggedIn ?? false)) {
            Navigator.pushNamed(context, "/login");
          }
          return (Scaffold(
            appBar: AppBar(
              title: Text(widget.title),
            ),
            body: Center(
              child: Column(
                mainAxisAlignment: MainAxisAlignment.center,
                children: <Widget>[
                  Text(
                    'Welcome ${session?.user?.username}',
                  ),
                ],
              ),
            ),
          ));
        });
  }
}
