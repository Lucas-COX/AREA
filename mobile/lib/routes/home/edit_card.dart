import 'package:flutter/material.dart';
import 'package:mobil/routes/home/services/service_google_sign_in.dart';
import 'package:mobil/routes/home/services/service_triggers.dart';
import '../../services/services_session.dart';
import '../login/login.dart';
import '../services/services/services_services.dart';

class EditCard extends StatefulWidget {
  const EditCard({Key? key, required this.title}) : super(key: key);
  final String title;

  @override
  State<EditCard> createState() => _EditCardState();
}

class _EditCardState extends State<EditCard> {
  List<String> get_list_reactions_name(List<String> list) {
    List<String> tmp = [];
    for (int i = 0; i < list.length; i++) {
      if (list[i] == 'microsoft' || list[i] == 'timer' || list[i] == 'google') {
        continue;
      } else {
        tmp.add(list[i]);
      }
    }
    tmp.add('Undefined');
    return tmp;
  }

  List<String> get_list_actions_name(List<String> list) {
    List<String> tmp = [];
    for (int i = 0; i < list.length; i++) {
      if (list[i] == 'discord' || list[i] == 'notion') {
        continue;
      } else {
        tmp.add(list[i]);
      }
    }
    tmp.add('Undefined');
    return tmp;
  }

  List<ServiceReaction> get_list_reactions_reactions(
      String naming, List<Service>? services) {
    if (services == null) {
      return [];
    }
    List<String> tmp = services.map((e) => e.name).toList();
    if (tmp.contains(naming)) {
      if (services.firstWhere((name) => name.name == naming).reactions != []) {
        print(services.firstWhere((name) => name.name == naming).reactions);
        return services.firstWhere((name) => name.name == naming).reactions;
      } else {
        return [];
      }
    }
    return [];
  }

  List<ServiceAction> get_list_actions_actions(
      String naming, List<Service>? services) {
    if (services == null) {
      return [];
    }
    List<String> tmp = services.map((e) => e.name).toList();

    if (tmp.contains(naming)) {
      if (services.firstWhere((name) => name.name == naming).actions != []) {
        return services.firstWhere((name) => name.name == naming).actions;
      } else {
        return [];
      }
    }
    return [];
  }

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
            return AccueilPage(title: widget.title);
          }
          final trigger = ModalRoute.of(context)!.settings.arguments as Trigger;
          return (FutureBuilder<List<Service>>(
              future: Services.getServices(),
              builder: (BuildContext context,
                  AsyncSnapshot<List<Service>> snapshot) {
                var services = snapshot.data;
                List<String>? servicesAvalable =
                    session?.user == null ? [] : session?.user?.services;
                services?.insert(services.length,
                    Service(name: "Undefined", actions: [], reactions: []));
                List<ServiceAction> listActions =
                    get_list_actions_actions(trigger.actionService, services);
                List<ServiceReaction> listReactions =
                    get_list_reactions_reactions(
                        trigger.reactionService, services);
                return (Scaffold(
                  backgroundColor: const Color.fromRGBO(255, 252, 242, 1),
                  appBar: AppBar(
                    backgroundColor: const Color.fromRGBO(255, 252, 242, 1),
                    title: const Text('Edit card',
                        style: TextStyle(
                            color: Color.fromRGBO(37, 36, 34, 1),
                            fontFamily: 'Roboto')),
                  ),
                  body: Center(
                      child: Column(
                    children: [
                      const Text('Edit card',
                          style: TextStyle(
                              color: Color.fromRGBO(37, 36, 34, 1),
                              fontFamily: 'Roboto')),
                      const SizedBox(height: 50),
                      Padding(
                          padding: const EdgeInsets.all(10),
                          child: TextFormField(
                            initialValue: trigger.title,
                            decoration: const InputDecoration(
                              focusedBorder: OutlineInputBorder(
                                borderSide: BorderSide(
                                    color: Color.fromRGBO(235, 94, 40, 1),
                                    width: 2.0),
                              ),
                              border: OutlineInputBorder(
                                  borderRadius:
                                      BorderRadius.all(Radius.circular(10.0)),
                                  borderSide: BorderSide(color: Colors.black)),
                              filled: true,
                              labelText: 'Title',
                              labelStyle: TextStyle(
                                  color: Color.fromRGBO(37, 36, 34, 1)),
                            ),
                            onChanged: (value) {
                              trigger.title = value.trim();
                            },
                          )),
                      Padding(
                          padding: const EdgeInsets.all(10),
                          child: TextFormField(
                            initialValue: trigger.description,
                            decoration: const InputDecoration(
                              focusedBorder: OutlineInputBorder(
                                borderSide: BorderSide(
                                    color: Color.fromRGBO(235, 94, 40, 1),
                                    width: 2.0),
                              ),
                              border: OutlineInputBorder(
                                  borderRadius:
                                      BorderRadius.all(Radius.circular(10.0)),
                                  borderSide: BorderSide(color: Colors.black)),
                              filled: true,
                              labelText: 'Description',
                              labelStyle: TextStyle(
                                  color: Color.fromRGBO(37, 36, 34, 1)),
                            ),
                            onChanged: (value) {
                              trigger.title = value.trim();
                            },
                          )),
                      const SizedBox(height: 50),
                      const Text('Action',
                          style: TextStyle(
                              color: Color.fromRGBO(37, 36, 34, 1),
                              fontFamily: 'Roboto')),
                      Card(
                        color: const Color.fromRGBO(235, 94, 40, 1),
                        shape: RoundedRectangleBorder(
                            borderRadius: BorderRadius.circular(10)),
                        child: Column(
                          children: <Widget>[
                            Row(children: <Widget>[
                              const SizedBox(width: 20),
                              DropdownButton<String>(
                                  value: trigger.actionService == ''
                                      ? 'Undefined'
                                      : trigger.actionService,
                                  icon: const Icon(Icons.arrow_downward),
                                  iconSize: 24,
                                  elevation: 16,
                                  style: const TextStyle(color: Colors.black),
                                  underline: Container(
                                    height: 2,
                                    color: Colors.black,
                                  ),
                                  onChanged: (String? newValue) {
                                    setState(() {
                                      trigger.actionService = newValue!;
                                      listActions = get_list_actions_actions(
                                          trigger.actionService, services);
                                    });
                                  },
                                  items:
                                      get_list_actions_name(servicesAvalable!)
                                          .map<DropdownMenuItem<String>>(
                                              (String value) {
                                    return DropdownMenuItem<String>(
                                      value: value,
                                      child: Text(value),
                                    );
                                  }).toList(),
                                  hint: const Text('Service')),
                              const SizedBox(width: 50),
                              if (listActions.isNotEmpty) ...[
                                Expanded(
                                    child: DropdownButton<String>(
                                  value: listActions.isNotEmpty
                                      ? listActions.first.name
                                      : 'Undefined',
                                  icon: const Icon(Icons.arrow_downward),
                                  iconSize: 24,
                                  elevation: 16,
                                  style: const TextStyle(color: Colors.black),
                                  underline: Container(
                                    height: 2,
                                    color: Colors.black,
                                  ),
                                  onChanged: (String? newValue) {
                                    setState(() {
                                      trigger.action = newValue!;
                                    });
                                  },
                                  items: get_list_actions_actions(
                                          trigger.actionService, services)
                                      .map<DropdownMenuItem<String>>(
                                          (ServiceAction value) {
                                    return DropdownMenuItem<String>(
                                      value: value.name,
                                      child: Text(value.name),
                                    );
                                  }).toList(),
                                ))
                              ],
                            ]),
                            if (trigger.actionService == 'github' ||
                                trigger.actionService == 'timer') ...[
                              const SizedBox(height: 10),
                            ],
                            Row(
                              children: <Widget>[
                                if (trigger.actionService == 'github') ...[
                                  Expanded(
                                    child: Padding(
                                      padding: const EdgeInsets.symmetric(
                                          horizontal: 20),
                                      child: TextFormField(
                                        obscureText: true,
                                        decoration: const InputDecoration(
                                          hoverColor:
                                              Color.fromRGBO(235, 94, 40, 1),
                                          labelText: 'Repository',
                                          contentPadding: EdgeInsets.all(14),
                                          focusedBorder: OutlineInputBorder(
                                            borderSide: BorderSide(
                                                color: Color.fromRGBO(
                                                    235, 94, 40, 1),
                                                width: 2.0),
                                            borderRadius: BorderRadius.all(
                                                Radius.circular(10)),
                                          ),
                                          border: OutlineInputBorder(
                                            borderRadius: BorderRadius.all(
                                                Radius.circular(10)),
                                            borderSide: BorderSide(
                                                color: Color.fromRGBO(
                                                    235, 94, 40, 1)),
                                          ),
                                        ),
                                        onChanged: (value) => setState(() {
                                          trigger.actionData = value;
                                        }),
                                      ),
                                    ),
                                  )
                                ],
                                if (trigger.actionService == 'timer') ...[
                                  Expanded(
                                      child: Padding(
                                          padding: const EdgeInsets.symmetric(
                                              horizontal: 20),
                                          child: TextFormField(
                                            obscureText: true,
                                            decoration: const InputDecoration(
                                              hoverColor: Color.fromRGBO(
                                                  235, 94, 40, 1),
                                              labelText: 'Time',
                                              contentPadding:
                                                  EdgeInsets.all(14),
                                              focusedBorder: OutlineInputBorder(
                                                borderSide: BorderSide(
                                                    color: Color.fromRGBO(
                                                        235, 94, 40, 1),
                                                    width: 2.0),
                                                borderRadius: BorderRadius.all(
                                                    Radius.circular(10)),
                                              ),
                                              border: OutlineInputBorder(
                                                borderRadius: BorderRadius.all(
                                                    Radius.circular(10)),
                                                borderSide: BorderSide(
                                                    color: Color.fromRGBO(
                                                        235, 94, 40, 1)),
                                              ),
                                            ),
                                            onChanged: (value) => setState(() {
                                              trigger.actionData = value;
                                            }),
                                          )))
                                ]
                              ],
                            ),
                            if (trigger.actionService == 'github' ||
                                trigger.actionService == 'timer') ...[
                              const SizedBox(height: 10),
                            ],
                          ],
                        ),
                      ),
                      const SizedBox(height: 50),
                      const Text('Reaction',
                          style: TextStyle(
                              color: Color.fromRGBO(37, 36, 34, 1),
                              fontFamily: 'Roboto')),
                      Card(
                        color: const Color.fromRGBO(235, 94, 40, 1),
                        shape: RoundedRectangleBorder(
                            borderRadius: BorderRadius.circular(10)),
                        child: Column(
                          children: <Widget>[
                            Row(children: <Widget>[
                              const SizedBox(width: 20),
                              DropdownButton<String>(
                                  value: trigger.reactionService == ''
                                      ? 'Undefined'
                                      : trigger.reactionService,
                                  icon: const Icon(Icons.arrow_downward),
                                  iconSize: 24,
                                  elevation: 16,
                                  style: const TextStyle(color: Colors.black),
                                  underline: Container(
                                    height: 2,
                                    color: Colors.black,
                                  ),
                                  onChanged: (String? newValue) {
                                    setState(() {
                                      trigger.reactionService = newValue!;
                                      listReactions =
                                          get_list_reactions_reactions(
                                              trigger.reactionService,
                                              services);
                                    });
                                  },
                                  items:
                                      get_list_reactions_name(servicesAvalable!)
                                          .map<DropdownMenuItem<String>>(
                                              (String value) {
                                    return DropdownMenuItem<String>(
                                      value: value,
                                      child: Text(value),
                                    );
                                  }).toList(),
                                  hint: const Text('Service')),
                              const SizedBox(width: 50),
                              if (listReactions.isNotEmpty) ...[
                                Expanded(
                                    child: DropdownButton<String>(
                                  value: listReactions.isNotEmpty
                                      ? listReactions.first.name
                                      : 'Undefined',
                                  icon: const Icon(Icons.arrow_downward),
                                  iconSize: 24,
                                  elevation: 16,
                                  style: const TextStyle(color: Colors.black),
                                  underline: Container(
                                    height: 2,
                                    color: Colors.black,
                                  ),
                                  onChanged: (String? newValue) {
                                    setState(() {
                                      trigger.reaction = newValue!;
                                    });
                                  },
                                  items: get_list_reactions_reactions(
                                          trigger.reactionService, services)
                                      .map<DropdownMenuItem<String>>(
                                          (ServiceReaction value) {
                                    return DropdownMenuItem<String>(
                                      value: value.name,
                                      child: Text(value.name),
                                    );
                                  }).toList(),
                                ))
                              ],
                            ]),
                            if (trigger.reactionService == 'discord') ...[
                              const SizedBox(height: 10),
                            ],
                            const SizedBox(height: 20),
                            Row(
                              children: <Widget>[
                                if (trigger.reactionService == 'discord') ...[
                                  Expanded(
                                    child: Padding(
                                      padding: const EdgeInsets.symmetric(
                                          horizontal: 20),
                                      child: TextFormField(
                                        obscureText: true,
                                        decoration: const InputDecoration(
                                          hoverColor:
                                              Color.fromRGBO(235, 94, 40, 1),
                                          labelText: 'Weebhook',
                                          contentPadding: EdgeInsets.all(14),
                                          focusedBorder: OutlineInputBorder(
                                            borderSide: BorderSide(
                                                color: Color.fromRGBO(
                                                    235, 94, 40, 1),
                                                width: 2.0),
                                            borderRadius: BorderRadius.all(
                                                Radius.circular(10)),
                                          ),
                                          border: OutlineInputBorder(
                                            borderRadius: BorderRadius.all(
                                                Radius.circular(10)),
                                            borderSide: BorderSide(
                                                color: Color.fromRGBO(
                                                    235, 94, 40, 1)),
                                          ),
                                        ),
                                        onChanged: (value) => setState(() {
                                          trigger.reactionData = value;
                                        }),
                                      ),
                                    ),
                                  )
                                ],
                                if (trigger.reactionService == 'github') ...[
                                  Expanded(
                                    child: Padding(
                                      padding: const EdgeInsets.symmetric(
                                          horizontal: 20),
                                      child: TextFormField(
                                        obscureText: true,
                                        decoration: const InputDecoration(
                                          hoverColor:
                                              Color.fromRGBO(235, 94, 40, 1),
                                          labelText: 'Repository',
                                          contentPadding: EdgeInsets.all(14),
                                          focusedBorder: OutlineInputBorder(
                                            borderSide: BorderSide(
                                                color: Color.fromRGBO(
                                                    235, 94, 40, 1),
                                                width: 2.0),
                                            borderRadius: BorderRadius.all(
                                                Radius.circular(10)),
                                          ),
                                          border: OutlineInputBorder(
                                            borderRadius: BorderRadius.all(
                                                Radius.circular(10)),
                                            borderSide: BorderSide(
                                                color: Color.fromRGBO(
                                                    235, 94, 40, 1)),
                                          ),
                                        ),
                                        onChanged: (value) => setState(() {
                                          trigger.reactionData = value;
                                        }),
                                      ),
                                    ),
                                  )
                                ],
                              ],
                            ),
                            if (trigger.reactionService == 'discord' ||
                                trigger.reactionService == 'github') ...[
                              const SizedBox(height: 10),
                            ],
                          ],
                        ),
                      ),
                      const SizedBox(height: 50),
                      ElevatedButton(
                          style: ElevatedButton.styleFrom(
                            foregroundColor:
                                const Color.fromRGBO(235, 94, 40, 1),
                            backgroundColor:
                                const Color.fromRGBO(255, 252, 242, 1),
                            shape: RoundedRectangleBorder(
                              borderRadius: BorderRadius.circular(10),
                            ),
                          ),
                          onPressed: (() async {
                            final response = await TriggersService.put(
                                trigger.toTriggerBody(), trigger.id);
                            if (response.statusCode == 200) {
                              Navigator.pop(context);
                            }
                          }),
                          child: const Text('Save')),
                    ],
                  )),
                ));
              }));
        });
  }
}

class ActionGoogle extends StatefulWidget {
  const ActionGoogle({Key? key}) : super(key: key);

  @override
  State<ActionGoogle> createState() => _ActionGoogleState();
}

class _ActionGoogleState extends State<ActionGoogle> {
  @override
  Widget build(BuildContext context) {
    return Card(
      color: const Color.fromRGBO(255, 252, 242, 1),
      child: Container(
        alignment: Alignment.center,
        height: 150,
        width: 150,
        decoration: BoxDecoration(
          borderRadius: BorderRadius.circular(10),
          border: Border.all(
            color: const Color.fromRGBO(235, 94, 40, 1),
          ),
        ),
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            const Text(
              'ServiceAction :',
              style: TextStyle(
                fontSize: 20,
                fontWeight: FontWeight.bold,
                color: Color.fromRGBO(37, 36, 34, 1),
              ),
            ),
            IconButton(
                onPressed: () async {
                  final url = await Openwindow.getUrl();
                  await Openwindow.openwindow(url['url']);
                },
                icon: const Icon(
                  Icons.g_mobiledata_rounded,
                  size: 50,
                  color: Color.fromRGBO(37, 36, 34, 1),
                )),
            const Text('Connect to google',
                style: TextStyle(
                  fontSize: 11,
                  color: Color.fromRGBO(37, 36, 34, 1),
                )),
          ],
        ),
      ),
    );
  }
}
