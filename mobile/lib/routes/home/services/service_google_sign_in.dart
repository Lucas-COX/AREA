import 'package:flutter/widgets.dart';
import 'package:url_launcher/url_launcher.dart';
import 'package:shared_preferences/shared_preferences.dart';
import 'dart:async';
import 'dart:convert';
import 'package:http/http.dart' as http;

class Openwindow {
  static Future getUrl() async {
    var completer = Completer();
    final prefs = await SharedPreferences.getInstance();
    final token = prefs.getString('area_token');
    Codec<String, String> stringToBase64Url = utf8.fuse(base64Url);
    String url =
        '${const String.fromEnvironment('API_URL')}/providers/google/auth?callback=${stringToBase64Url.encode('https://google.com')}';
    if (token != null) {
      try {
        final response =
            await http.get(Uri.parse(url), headers: <String, String>{
          'Authorization': 'Bearer $token',
        });
        completer.complete(jsonDecode(response.body));
      } catch (e) {
        debugPrint(e.toString());
      }
      return completer.future;
    }
  }

  static Future openwindow(String url) async {
    if (!await launchUrl(Uri.parse(url),
        mode: LaunchMode.externalApplication)) {
      throw 'Could not launch $url';
    }
  }
}
