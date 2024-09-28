import 'package:google_sign_in/google_sign_in.dart';

class LoginApi {
  static final _googleSignIn = GoogleSignIn(
    clientId:
        "1022719781194-udr6ghqd929sv086vb60n66lnkjds3lb.apps.googleusercontent.com",
  );
  static Future<GoogleSignInAccount?> login() => _googleSignIn.signIn();
  static Future signOut = _googleSignIn.signOut();
}
