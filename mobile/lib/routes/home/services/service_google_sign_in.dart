import 'package:google_sign_in/google_sign_in.dart';

final GoogleSignIn googleSignIn = GoogleSignIn(
  scopes: <String>[
    'email',
    'https://www.googleapis.com/auth/contacts.readonly',
  ],
  clientId:
      '613091759382-6tq4hm6et3iad5u70hkk96imjvv0ee6d.apps.googleusercontent.com',
  serverClientId:
      '613091759382-6hjja3qv3866thps1mru6v139pti64ju.apps.googleusercontent.com',
  forceCodeForRefreshToken: true,
);

Future<void> handleSignIn() async {
  try {
    await googleSignIn.signIn();
  } catch (error) {
    print(error);
  }
}
