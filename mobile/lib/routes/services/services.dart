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
  Map<String, Icon> icons = {
    'google': const Icon(Icons.mail),
    'discord': const Icon(Icons.discord),
    'microsoft': const Icon(Icons.mail),
    'github': const Icon(Icons.code),
    'notion': const Icon(Icons.note),
    'timer': const Icon(Icons.timer),
    'default': const Icon(Icons.question_mark),
  };

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
                    backgroundColor: const Color.fromRGBO(255, 252, 242, 1),
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
                              color: session?.user?.services
                                          .contains(services?[index].name) ==
                                      true
                                  ? const Color.fromRGBO(235, 94, 40, 1)
                                  : const Color.fromRGBO(255, 252, 242, 1),
                              child: ListTile(
                                  title: Text(
                                    services?[index].name ?? '',
                                    style: TextStyle(
                                      color: session?.user?.services.contains(
                                                  services?[index].name) ==
                                              true
                                          ? const Color.fromRGBO(
                                              255, 252, 242, 1)
                                          : const Color.fromRGBO(37, 36, 34, 1),
                                    ),
                                  ),
                                  leading:
                                      icons[services?[index].name ?? 'default'],
                                  onTap: () async {
                                    if (services?[index].name == 'discord' ||
                                        services?[index].name == 'timer') {
                                      await Services.getUrl(
                                          services?[index].name ?? '');
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
