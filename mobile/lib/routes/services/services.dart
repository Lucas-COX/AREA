import 'package:flutter/material.dart';
import '../../services/services_session.dart';
import '../login/login.dart';
import 'services/services_services.dart';

class ServicesPage extends StatefulWidget {
  const ServicesPage({Key? key, required this.title}) : super(key: key);
  final String title;

  @override
  State<ServicesPage> createState() => _ServicesPageState();
}

class _ServicesPageState extends State<ServicesPage> {
  get_icon(String name) {
    if (name == 'gmail') {
      return const Icon(Icons.mail);
    }
    if (name == 'discord') {
      return const Icon(Icons.discord);
    }
    if (name == 'outlook') {
      return const Icon(Icons.mail);
    }
    if (name == 'github') {
      return const Icon(Icons.code);
    }
    if (name == 'notion') {
      return const Icon(Icons.note);
    }
  }

  @override
  Widget build(BuildContext context) {
    return FutureBuilder<Session>(
        future: ServicesSession.get(),
        builder: (BuildContext context, AsyncSnapshot<Session> snapshot) {
          if (!snapshot.hasData) {
            return (const Scaffold(
                body: Center(
                    child: CircularProgressIndicator(
              color: Colors.amberAccent,
            ))));
          }
          final session = snapshot.data;
          if (session?.isLoggedIn != null && !(session?.isLoggedIn ?? false)) {
            return AccueilPage(title: widget.title);
          }
          return (FutureBuilder<List<Service>>(
              future: Services.getServices(),
              builder: (BuildContext context,
                  AsyncSnapshot<List<Service>> snapshot) {
                if (!snapshot.hasData) {
                  return (const Scaffold(
                      body: Center(
                          child: CircularProgressIndicator(
                    color: Colors.amberAccent,
                  ))));
                }
                final services = snapshot.data;
                return (Scaffold(
                    appBar: AppBar(
                      centerTitle: true,
                      title: Text(widget.title,
                          style: const TextStyle(color: Colors.black)),
                      backgroundColor: const Color.fromRGBO(255, 252, 242, 1),
                      elevation: 0,
                    ),
                    body: ListView.builder(
                        itemCount: services?.length,
                        itemBuilder: (BuildContext context, int index) {
                          return (Card(
                              color: const Color.fromRGBO(235, 94, 40, 1),
                              child: ListTile(
                                  title: Text(
                                    services?[index].name ?? '',
                                    style: const TextStyle(
                                      color: Colors.white,
                                    ),
                                  ),
                                  leading:
                                      get_icon(services?[index].name ?? ''),
                                  onTap: () async {
                                    String weebok = '';
                                    if (services?[index].name == 'discord') {
                                      showDialog(
                                          context: context,
                                          builder: (context) {
                                            return AlertDialog(
                                              title: const Text('Discord'),
                                              content: TextFormField(
                                                initialValue: '',
                                                decoration:
                                                    const InputDecoration(
                                                        border:
                                                            OutlineInputBorder(),
                                                        labelText: 'Weebhook'),
                                                onChanged: (value) {
                                                  weebok = value;
                                                },
                                              ),
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
                                    } else {
                                      final url = await Services.getUrl(
                                          services?[index].name ?? '');
                                      await Services.connexion(url['url']);
                                    }
                                  })));
                        })));
              }));
        });
  }
}
