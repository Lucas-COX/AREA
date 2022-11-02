import 'package:flutter/material.dart';
import 'package:mobil/routes/home/services/service_google_sign_in.dart';
import 'package:mobil/routes/home/services/service_triggers.dart';
import '../../services/services_session.dart';
import '../login/login.dart';

class EditCard extends StatefulWidget {
  const EditCard({Key? key, required this.title}) : super(key: key);
  final String title;

  @override
  State<EditCard> createState() => _EditCardState();
}

class _EditCardState extends State<EditCard> {
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
          return (Scaffold(
              backgroundColor: const Color.fromRGBO(255, 252, 242, 1),
              appBar: AppBar(
                title: const Text('Edit your trigger'),
                centerTitle: true,
                backgroundColor: const Color.fromRGBO(255, 252, 242, 1),
              ),
              body: Form(
                child: Column(
                  mainAxisAlignment: MainAxisAlignment.start,
                  children: [
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
                            labelStyle:
                                TextStyle(color: Color.fromRGBO(37, 36, 34, 1)),
                          ),
                          onChanged: (value) {
                            trigger.title = value.trim();
                          },
                        )),
                    const SizedBox(height: 50),
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
                            labelStyle:
                                TextStyle(color: Color.fromRGBO(37, 36, 34, 1)),
                          ),
                          onChanged: (value) {
                            trigger.description = value.trim();
                          },
                        )),
                    const SizedBox(height: 50),
                    const SizedBox(height: 50),
                    ElevatedButton(
                        style: ElevatedButton.styleFrom(
                          foregroundColor: const Color.fromRGBO(235, 94, 40, 1),
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
                            Navigator.pop(context, '/');
                          }
                        }),
                        child: const Text('Save')),
                  ],
                ),
              )));
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
              'Action :',
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

//  Card(
//                       color: const Color.fromRGBO(235, 94, 40, 1),
//                       child: Container(
//                         margin: const EdgeInsets.all(10),
//                         decoration: BoxDecoration(
//                           borderRadius: BorderRadius.circular(10),
//                           border: Border.all(
//                             color: const Color.fromRGBO(235, 94, 40, 1),
//                           ),
//                         ),
//                         child: Row(
//                           mainAxisAlignment: MainAxisAlignment.spaceEvenly,
//                           children: [
//                             Padding(
//                               padding: const EdgeInsets.all(5),
//                               child: Row(
//                                 mainAxisAlignment:
//                                     MainAxisAlignment.spaceEvenly,
//                                 children: [
//                                   const ActionGoogle(),
//                                   const Icon(
//                                     Icons.arrow_right_alt_rounded,
//                                     size: 35,
//                                     color: Color.fromRGBO(37, 36, 34, 1),
//                                   ),
//                                   Card(
//                                     color:
//                                         const Color.fromRGBO(255, 252, 242, 1),
//                                     child: Container(
//                                       alignment: Alignment.center,
//                                       height: 150,
//                                       width: 150,
//                                       decoration: BoxDecoration(
//                                         borderRadius: BorderRadius.circular(10),
//                                         border: Border.all(
//                                           color: const Color.fromRGBO(
//                                               235, 94, 40, 1),
//                                         ),
//                                       ),
//                                       child: Column(
//                                         mainAxisAlignment:
//                                             MainAxisAlignment.center,
//                                         children: [
//                                           const Text(
//                                             'Reaction :',
//                                             style: TextStyle(
//                                               fontSize: 20,
//                                               fontWeight: FontWeight.bold,
//                                               color:
//                                                   Color.fromRGBO(37, 36, 34, 1),
//                                             ),
//                                           ),
//                                           Padding(
//                                               padding: const EdgeInsets.all(10),
//                                               child: TextFormField(
//                                                 initialValue:
//                                                     trigger.reactionData,
//                                                 decoration:
//                                                     const InputDecoration(
//                                                   focusedBorder:
//                                                       OutlineInputBorder(
//                                                     borderRadius:
//                                                         BorderRadius.all(
//                                                             Radius.circular(
//                                                                 10.0)),
//                                                     borderSide: BorderSide(
//                                                         color: Color.fromRGBO(
//                                                             235, 94, 40, 1),
//                                                         width: 1.0),
//                                                   ),
//                                                   border: OutlineInputBorder(
//                                                       borderRadius:
//                                                           BorderRadius.all(
//                                                               Radius.circular(
//                                                                   10.0)),
//                                                       borderSide: BorderSide(
//                                                           color: Colors.black)),
//                                                   filled: true,
//                                                   labelText:
//                                                       'Weebhook for discord',
//                                                   labelStyle: TextStyle(
//                                                       fontSize: 10,
//                                                       color: Color.fromRGBO(
//                                                           37, 36, 34, 1)),
//                                                 ),
//                                                 onChanged: (value) {
//                                                   trigger.reactionData =
//                                                       value.trim();
//                                                 },
//                                               )),
//                                           const Text('Connect to discord url',
//                                               style: TextStyle(
//                                                 fontSize: 11,
//                                                 color: Color.fromRGBO(
//                                                     37, 36, 34, 1),
//                                               )),
//                                         ],
//                                       ),
//                                     ),
//                                   )
//                                 ],
//                               ),
//                             ),
//                           ],
//                         ),
//                       ),
//                     ),