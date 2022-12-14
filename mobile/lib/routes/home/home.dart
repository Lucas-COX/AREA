import 'package:draggable_home/draggable_home.dart';
import 'package:flutter/material.dart';
import 'package:mobil/routes/login/login.dart';
import 'services/service_triggers.dart';

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
                body: Center(
                    child: CircularProgressIndicator(
              color: Colors.amberAccent,
            ))));
          }
          final session = snapshot.data;
          if (session?.isLoggedIn != null && !(session?.isLoggedIn ?? false)) {
            return AccueilPage(title: widget.title);
          }
          final triggers = session?.user == null ? [] : session?.user?.triggers;
          return (DraggableHome(
              title: const Text("",
                  style: TextStyle(color: Color.fromRGBO(206, 13, 13, 1))),
              backgroundColor: const Color.fromRGBO(255, 255, 255, 1),
              headerWidget: headerWidget(context, session),
              alwaysShowTitle: true,
              alwaysShowLeadingAndAction: true,
              curvedBodyRadius: 30,
              leading: IconButton(
                  onPressed: (() {
                    Navigator.pushNamed(context, '/login');
                  }),
                  icon: const Icon(Icons.login,
                      color: Color.fromRGBO(206, 13, 13, 1))),
              appBarColor: const Color.fromRGBO(255, 255, 255, 1),
              body: [
                const Text(
                  'Your trigger',
                  style: TextStyle(
                    fontSize: 20,
                    fontWeight: FontWeight.bold,
                    color: Color.fromRGBO(206, 13, 13, 1),
                  ),
                ),
                const SizedBox(height: 50),
                ListView.builder(
                    padding: const EdgeInsets.only(top: 0),
                    physics: const NeverScrollableScrollPhysics(),
                    shrinkWrap: true,
                    itemCount: triggers?.length ?? 0,
                    itemBuilder: (context, index) => Card(
                          color: const Color.fromRGBO(12, 169, 12, 1),
                          child: ListTile(
                            subtitle: Text(triggers?[index].description ??
                                'No description'),
                            title: Text(triggers?[index].title ?? ''),
                            trailing: Row(
                              mainAxisSize: MainAxisSize.min,
                              children: [
                                Switch(
                                  value: triggers?[index].active,
                                  activeColor: Color.fromRGBO(206, 13, 13, 1),
                                  onChanged: (value) async {
                                    setState(() {
                                      print(value);
                                      print(triggers?[index].active);
                                      triggers?[index].active = value;
                                      print(triggers?[index].active);
                                    });
                                    await TriggersService.update(
                                        triggers?[index].toTriggerBody(),
                                        triggers?[index].id);
                                  },
                                ),
                                IconButton(
                                    onPressed: () async {
                                      await TriggersService.delete(
                                          triggers?[index].id);
                                      setState(() {
                                        triggers?.removeAt(index);
                                      });
                                    },
                                    icon: const Icon(Icons.delete,
                                        color: Color.fromRGBO(206, 13, 13, 1))),
                                IconButton(
                                    onPressed: () async {
                                      Navigator.pushNamed(context, '/edit',
                                          arguments: triggers?[index]);
                                    },
                                    icon: const Icon(
                                      Icons.edit,
                                      color: Color.fromRGBO(206, 13, 13, 1),
                                    )),
                              ],
                            ),
                            contentPadding: const EdgeInsets.all(10),
                          ),
                        ))
              ]));
        });
  }
}

Widget headerWidget(BuildContext context, final session) {
  return Container(
      height: 200,
      width: double.infinity,
      color: const Color.fromRGBO(255, 255, 255, 1),
      child:
          Column(mainAxisAlignment: MainAxisAlignment.start, children: <Widget>[
        const SizedBox(height: 50),
        Text(
          'Welcome ${session?.user?.username}',
          style: const TextStyle(
            fontSize: 30,
            fontWeight: FontWeight.bold,
            color: Color.fromRGBO(206, 13, 13, 1),
          ),
        ),
        const SizedBox(height: 50),
        ElevatedButton(
            style: ElevatedButton.styleFrom(
                primary: const Color.fromRGBO(12, 169, 12, 1)),
            onPressed: () {
              Navigator.pushNamed(context, '/services');
            },
            child: const Text('Connect to your services',
                style: TextStyle(color: Colors.white))),
        const SizedBox(height: 10),
        const Text(
          'Create your trigger',
          style: TextStyle(
            fontSize: 20,
            fontWeight: FontWeight.bold,
            color: Color.fromRGBO(206, 13, 13, 1),
          ),
        ),
        const SizedBox(height: 10),
        ElevatedButton(
          style: ElevatedButton.styleFrom(
            foregroundColor: Colors.white,
            backgroundColor: const Color.fromRGBO(12, 169, 12, 1),
            shape: RoundedRectangleBorder(
              borderRadius: BorderRadius.circular(32.0),
            ),
            padding: const EdgeInsets.symmetric(horizontal: 50, vertical: 15),
          ),
          onPressed: () async {
            final response = await TriggersService.post();
            if (response.statusCode == 200) {
              Navigator.pushNamed(context, '/');
            }
          },
          child: const Icon(Icons.add),
        ),
      ]));
}
