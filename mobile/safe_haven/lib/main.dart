import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
import 'package:safe_haven/features/auth/presentation/bloc/bloc/auth_bloc_bloc.dart';
import 'package:safe_haven/features/auth/presentation/screens/forgot_password.dart';
import 'package:safe_haven/features/auth/presentation/screens/log_in.dart';
import 'package:safe_haven/features/auth/presentation/screens/reset_password.dart';
import 'package:safe_haven/features/auth/presentation/screens/sign_up.dart';
import 'package:safe_haven/features/auth/presentation/screens/sign_up_Phone.dart';
import 'package:safe_haven/features/case/presentation/bloc/case_bloc.dart';
import 'package:safe_haven/features/case/presentation/screens/create_case_screen.dart';
import 'package:safe_haven/injection_container.dart' as di;

import 'package:flutter_bloc/flutter_bloc.dart';
import 'injection_container.dart';

void main() async {
  WidgetsFlutterBinding.ensureInitialized();
  await di.init();
  di.sl<http.Client>();
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MultiBlocProvider(
      providers: [
        BlocProvider<AuthBlocBloc>(
          create: (context) => AuthBlocBloc(sl(), sl(), sl(), sl(), sl(), sl()),
        ),
        BlocProvider<CaseBloc>(
          create: (context) => CaseBloc(
            sl(),
          ),
        ),
        // Add more BlocProviders here as needed
      ],
      // Ensure Bloc is created here
      child: MaterialApp(
        debugShowCheckedModeBanner: false,
        title: 'Flutter Login',
        initialRoute: '/signup',
        onGenerateRoute: (settings) {
          switch (settings.name) {
            case '/login':
              return PageRouteBuilder(
                  pageBuilder: (context, _, __) => const LogInscreen());
            case '/signup':
              return PageRouteBuilder(
                  pageBuilder: (context, _, __) => const SignUpscreen());
            case '/forgotpassword':
              return PageRouteBuilder(
                  pageBuilder: (context, _, __) =>
                      const ForgotPasswordscreen());
            case '/resetpassword':
              return PageRouteBuilder(
                  pageBuilder: (context, _, __) => const ResetPasswordscreen());
            case '/signupphone':
              return PageRouteBuilder(
                  pageBuilder: (context, _, __) => const SignUpPhonescreen());
            case '/report':
              return PageRouteBuilder(
                  pageBuilder: (context, _, __) => CreateCaseScreen());

            // case '/checkvideo':
            //   return PageRouteBuilder(
            //       pageBuilder: (context, _, __) => VideoPickerScreen());
            default:
              return null; // Handle invalid routes
          }
        },
      ),
    );
  }
}

// class MyApp2 extends StatelessWidget {
//   @override
//   Widget build(BuildContext context) {
//     return MaterialApp(
//       title: 'Google Sign-In Demo',
//       home: SignInPage(),
//     );
//   }
// }

// class SignInPage extends StatefulWidget {
//   @override
//   _SignInPageState createState() => _SignInPageState();
// }

// class _SignInPageState extends State<SignInPage> {
//   final GoogleSignIn googleSignIn = GoogleSignIn(
//     serverClientId:
//         "1022719781194-udr6ghqd929sv086vb60n66lnkjds3lb.apps.googleusercontent.com",
//   );

//   GoogleSignInAccount? _user;

//   Future<void> _handleSignIn() async {
//     try {
//       print('ezi dersual omgggggggggggggggggggggggg');
//       _user = await googleSignIn.signIn();
//       print(_user!.displayName);
//       print(_user!.email);
//       // Handle successful sign-in
//       print('User signed in: ${_user?.email}');
//     } catch (error) {
//       print('Sign-in failed: $error');
//     }
//   }

//   Future<void> _handleSignOut() async {
//     await googleSignIn.signOut();
//     print('User signed out');
//   }

//   @override
//   Widget build(BuildContext context) {
//     return Scaffold(
//       appBar: AppBar(
//         title: Text("Google Sign-In"),
//       ),
//       body: Center(
//         child: Column(
//           mainAxisAlignment: MainAxisAlignment.center,
//           children: <Widget>[
//             if (_user != null)
//               Column(
//                 children: [
//                   Text('Signed in as: ${_user?.email}'),
//                   ElevatedButton(
//                     onPressed: _handleSignOut,
//                     child: Text('Sign Out'),
//                   ),
//                 ],
//               )
//             else
//               ElevatedButton(
//                 onPressed: _handleSignIn,
//                 child: Text('Sign In with Google'),
//               ),
//           ],
//         ),
//       ),
//     );
//   }
// }

// class GoogleSignInPage extends StatefulWidget {
//   @override
//   _GoogleSignInPageState createState() => _GoogleSignInPageState();
// }

// class _GoogleSignInPageState extends State<GoogleSignInPage> {
//   @override
//   Widget build(BuildContext context) {
//     return Scaffold(
//       appBar: AppBar(
//         title: const Text('loginGoogle'),
//       ),
//       body: SizedBox(
//         width: double.infinity,
//         child: Column(
//           mainAxisAlignment: MainAxisAlignment.center,
//           children: [
//             ElevatedButton(
//                 onPressed: () async {
//                   var user = await LoginApi.login();
//                   if (user != null) {
//                     print('ok sign in');
//                     print(user.displayName);
//                     print(user.email);
//                   }
//                 },
//                 child: Text('login google'))
//           ],
//         ),
//       ),
//     );
//   }
// }
