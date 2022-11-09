<<<<<<< HEAD
import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'routes/register/register.dart';
import 'routes/login/login.dart';
import 'routes/home/home.dart';
=======
import 'package:flutter/material.dart';
import 'routes/register/register.dart';
import 'routes/login/login.dart';
import 'routes/home/home.dart';
import 'routes/home/edit_card.dart';
import 'routes/services/services.dart';
>>>>>>> 8dc2ef7 (feat(mobile): creation of a functional flutter client (#72))

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      debugShowCheckedModeBanner: false,
      title: 'Flutter Demo',
      theme: ThemeData(
        useMaterial3: true,
      ),
      // do the route for the app
      initialRoute: '/',
      routes: {
        '/': (context) => const HomePage(
              title: 'Area',
            ),
        '/login': (context) => const AccueilPage(
              title: 'Area',
            ),
        '/register': (context) => const RegisterPage(),
<<<<<<< HEAD
=======
        '/edit': (context) => const EditCard(
              title: '',
            ),
        '/services': (context) => const ServicesPage(
              title: 'Services',
            ),
>>>>>>> 8dc2ef7 (feat(mobile): creation of a functional flutter client (#72))
      },
    );
  }
}
