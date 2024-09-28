// import 'dart:convert';
// import 'package:flutter/gestures.dart';
import 'package:flutter/material.dart';
// import 'package:flutter_svg/svg.dart';
import 'package:http/http.dart' as http;
// import 'package:safe_haven/features/auth/domain/entities/sign_up_entity.dart';
import 'package:safe_haven/features/auth/presentation/bloc/bloc/auth_bloc_bloc.dart';
import 'package:safe_haven/features/auth/presentation/screens/forgot_password.dart';
import 'package:safe_haven/features/auth/presentation/screens/log_in.dart';
import 'package:safe_haven/features/auth/presentation/screens/reset_password.dart';
import 'package:safe_haven/features/auth/presentation/screens/sign_up.dart';
// import 'package:safe_haven/features/auth/presentation/widgets/custom_button.dart';
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
    return BlocProvider(
      create: (context) =>
          AuthBlocBloc(sl(), sl(), sl(), sl(), sl(), sl()), // Ensure Bloc is created here
      child: MaterialApp(
        debugShowCheckedModeBanner: false,
        title: 'Flutter Login',
        initialRoute: '/signup',
        onGenerateRoute: (settings) {
          switch (settings.name) {
            case '/login':
              return PageRouteBuilder(
                  pageBuilder: (context, _, __) => LogInscreen());
            case '/signup':
              return PageRouteBuilder(
                  pageBuilder: (context, _, __) => SignUpscreen());
            case '/forgotpassword':
              return PageRouteBuilder(
                  pageBuilder: (context, _, __) => ForgotPasswordscreen());
            case '/resetpassword':
              return PageRouteBuilder(
                  pageBuilder: (context, _, __) => ResetPasswordscreen());
            default:
              return null; // Handle invalid routes
          }
        },
      ),
    );
  }
}

// class SignUpscreen extends StatefulWidget {
//   const SignUpscreen({super.key});

//   @override
//   State<SignUpscreen> createState() => _SignUpScreen();
// }

// class _SignUpScreen extends State<SignUpscreen> {
//   // Controllers for form fields
//   final TextEditingController fullName = TextEditingController();
//   final TextEditingController email = TextEditingController();
//   final TextEditingController password = TextEditingController();
//   final TextEditingController confirmPassword = TextEditingController();

//   @override
//   Widget build(BuildContext context) {
//     return BlocListener<AuthBlocBloc, AuthBlocState>(
//       listener: (context, state) {
//         if (state is AuthRegisterSuccess) {
//           ScaffoldMessenger.of(context).showSnackBar(
//             SnackBar(
//               content: Text('Account created successfully!'),
//               backgroundColor: Theme.of(context).primaryColor,
//             ),
//           );
//           Navigator.pushNamed(context, '/login');
//         } else if (state is AuthError) {
//           ScaffoldMessenger.of(context).showSnackBar(
//             SnackBar(
//               content: Text(state.message),
//               backgroundColor: Colors.red,
//             ),
//           );
//         }
//       },
//       child: Scaffold(
//         appBar: AppBar(
//           title: Center(
//             child: Text('Register',
//                 style: TextStyle(color: Color(0xFF169C89), fontSize: 30)),
//           ),
//         ),
//         body: SingleChildScrollView(
//           padding: EdgeInsets.all(30),
//           child: Column(
//             children: [
//               CustomForm(
//                 fullName2: fullName,
//                 email2: email,
//                 password2: password,
//                 confirmPassword2: confirmPassword,
//               ),
//               SizedBox(height: 20),
//               BlocBuilder<AuthBlocBloc, AuthBlocState>(
//                 builder: (context, state) {
//                   return CustomButton(
//                     text: 'Register',
//                     onPressed: () {
//                       // Dispatch the registration event
//                       final customFormState =
//                           context.findAncestorStateOfType<_CustomFormState>();
//                       if (customFormState != null) {
//                         final selectedLanguage =
//                             customFormState.selectedLanguage;
//                         final selectedCategory =
//                             customFormState.selectedCategory;

//                         context.read<AuthBlocBloc>().add(RegisterEvent(
//                               registrationEntity: SignUpEntity(
//                                 language: selectedLanguage,
//                                 category: selectedCategory,
//                                 userType: 'normal',
//                                 fullName: fullName.text,
//                                 password: password.text,
//                               ),
//                             ));
//                       }
//                     },
//                     bC: 0xFFFFFFFF,
//                     col: 0xFF169C89,
//                   );
//                 },
//               ),
//               SizedBox(height: 20),
//               Row(
//                 mainAxisAlignment: MainAxisAlignment.center,
//                 children: [
//                   const Text('Have an account? '),
//                   RichText(
//                     text: TextSpan(
//                       text: 'Log In',
//                       style: TextStyle(color: Color(0xFF169C89)),
//                       recognizer: TapGestureRecognizer()
//                         ..onTap = () {
//                           Navigator.pushNamed(context, '/login');
//                         },
//                     ),
//                   ),
//                 ],
//               ),
//             ],
//           ),
//         ),
//       ),
//     );
//   }
// }

// class CustomForm extends StatefulWidget {
//   final TextEditingController fullName2;
//   final TextEditingController password2;
//   final TextEditingController confirmPassword2;
//   final TextEditingController email2;

//   const CustomForm({
//     super.key,
//     required this.fullName2,
//     required this.password2,
//     required this.confirmPassword2,
//     required this.email2,
//   });

//   @override
//   State<CustomForm> createState() => _CustomFormState();
// }

// class _CustomFormState extends State<CustomForm> {
//   // State variables to track the dropdown values
//   String _selectedLanguage = 'English'; // default value for language
//   String _selectedCategory = 'Victim'; // default value for category

//   String get selectedLanguage => _selectedLanguage;
//   String get selectedCategory => _selectedCategory;

//   @override
//   Widget build(BuildContext context) {
//     return SingleChildScrollView(
//       child: Padding(
//         padding: EdgeInsets.fromLTRB(10, 50, 10, 30),
//         child: Form(
//           child: Column(
//             children: [
//               // Full Name Field
//               Padding(
//                 padding: EdgeInsets.symmetric(horizontal: 10),
//                 child: const Align(
//                     alignment: Alignment.centerLeft,
//                     child: Row(
//                       children: [
//                         Text(
//                           'Full Name',
//                           style: TextStyle(
//                               fontSize: 15, fontWeight: FontWeight.w500),
//                         ),
//                         SizedBox(
//                           width: 5,
//                         ),
//                         Text(
//                           '*',
//                           style: TextStyle(
//                               fontSize: 15,
//                               fontWeight: FontWeight.w500,
//                               color: Colors.red),
//                         ),
//                       ],
//                     )),
//               ),
//               SizedBox(
//                 height: 15,
//               ),
//               TextFormField(
//                 controller: widget.fullName2,
//                 decoration: InputDecoration(
//                   hintText: 'Enter your name',
//                   hintStyle: TextStyle(color: Color(0xFFC7C7C7)),
//                   prefixIcon: Padding(
//                     padding: const EdgeInsets.all(12.0),
//                     child: Icon(Icons.person, color: Colors.grey),
//                   ),
//                   border: OutlineInputBorder(
//                     borderRadius: BorderRadius.all(Radius.circular(15.0)),
//                     borderSide: BorderSide(color: Color(0xFF169C89)),
//                   ),
//                   enabledBorder: OutlineInputBorder(
//                     borderRadius: BorderRadius.all(Radius.circular(15.0)),
//                     borderSide: BorderSide(color: Color(0xFF169C89)),
//                   ),
//                   focusedBorder: OutlineInputBorder(
//                     borderRadius: BorderRadius.all(Radius.circular(15.0)),
//                     borderSide: BorderSide(color: Color(0xFF169C89), width: 2),
//                   ),
//                   filled: true,
//                   fillColor: Color.fromARGB(255, 247, 245, 245),
//                 ),
//               ),

//               SizedBox(height: 20),

//               Padding(
//                 padding: EdgeInsets.symmetric(horizontal: 10),
//                 child: const Align(
//                     alignment: Alignment.centerLeft,
//                     child: Row(
//                       children: [
//                         Text(
//                           'Email address',
//                           style: TextStyle(
//                               fontSize: 15, fontWeight: FontWeight.w500),
//                         ),
//                         SizedBox(
//                           width: 5,
//                         ),
//                         Text(
//                           '*',
//                           style: TextStyle(
//                               fontSize: 15,
//                               fontWeight: FontWeight.w500,
//                               color: Colors.red),
//                         ),
//                       ],
//                     )),
//               ),
//               SizedBox(
//                 height: 15,
//               ),
//               TextFormField(
//                 controller: widget.email2,
//                 decoration: InputDecoration(
//                   hintText: 'Enter your email',
//                   hintStyle: TextStyle(color: Color(0xFFC7C7C7)),
//                   prefixIcon: Padding(
//                       padding: const EdgeInsets.all(12.0),
//                       child: Icon(
//                         Icons.email,
//                         color: Colors.grey,
//                       )),
//                   border: OutlineInputBorder(
//                     borderRadius: BorderRadius.all(Radius.circular(15.0)),
//                     borderSide: BorderSide(color: Color(0xFF169C89)),
//                   ),
//                   enabledBorder: OutlineInputBorder(
//                     borderRadius: BorderRadius.all(Radius.circular(15.0)),
//                     borderSide: BorderSide(color: Color(0xFF169C89)),
//                   ),
//                   focusedBorder: OutlineInputBorder(
//                     borderRadius: BorderRadius.all(Radius.circular(15.0)),
//                     borderSide: BorderSide(color: Color(0xFF169C89), width: 2),
//                   ),
//                   filled: true,
//                   fillColor: Color.fromARGB(255, 247, 245, 245),
//                 ),
//               ),
//               SizedBox(height: 20),
//               Padding(
//                 padding: EdgeInsets.symmetric(horizontal: 10),
//                 child: const Align(
//                     alignment: Alignment.centerLeft,
//                     child: Row(
//                       children: [
//                         Text(
//                           'Password',
//                           style: TextStyle(
//                               fontSize: 15, fontWeight: FontWeight.w500),
//                         ),
//                         SizedBox(
//                           width: 5,
//                         ),
//                         Text(
//                           '*',
//                           style: TextStyle(
//                               fontSize: 15,
//                               fontWeight: FontWeight.w500,
//                               color: Colors.red),
//                         ),
//                       ],
//                     )),
//               ),
//               SizedBox(
//                 height: 15,
//               ),

//               // Password Field
//               TextFormField(
//                 controller: widget.password2,
//                 decoration: InputDecoration(
//                   hintText: 'Enter your password',
//                   hintStyle: TextStyle(color: Color(0xFFC7C7C7)),
//                   prefixIcon: Padding(
//                     padding: const EdgeInsets.all(12.0),
//                     child: Icon(Icons.lock, color: Colors.grey),
//                   ),
//                   suffixIcon: Padding(
//                     padding: const EdgeInsets.all(12.0),
//                     child: SvgPicture.asset(
//                       'assets/icons/mdi_hide-outline.svg', // Your SVG file path
//                       width: 24, // Adjust width and height as needed
//                       height: 24,
//                     ),
//                   ),
//                   border: OutlineInputBorder(
//                     borderRadius: BorderRadius.all(Radius.circular(15.0)),
//                     borderSide: BorderSide(color: Color(0xFF169C89)),
//                   ),
//                   enabledBorder: OutlineInputBorder(
//                     borderRadius: BorderRadius.all(Radius.circular(15.0)),
//                     borderSide: BorderSide(color: Color(0xFF169C89)),
//                   ),
//                   focusedBorder: OutlineInputBorder(
//                     borderRadius: BorderRadius.all(Radius.circular(15.0)),
//                     borderSide: BorderSide(color: Color(0xFF169C89), width: 2),
//                   ),
//                   filled: true,
//                   fillColor: Color.fromARGB(255, 247, 245, 245),
//                 ),
//                 obscureText: true, // to hide password input
//               ),
//               SizedBox(height: 20),

//               // Confirm Password Field
//               Padding(
//                 padding: EdgeInsets.symmetric(horizontal: 10),
//                 child: const Align(
//                     alignment: Alignment.centerLeft,
//                     child: Row(
//                       children: [
//                         Text(
//                           'Confirm Password',
//                           style: TextStyle(
//                               fontSize: 15, fontWeight: FontWeight.w500),
//                         ),
//                         SizedBox(
//                           width: 5,
//                         ),
//                         Text(
//                           '*',
//                           style: TextStyle(
//                               fontSize: 15,
//                               fontWeight: FontWeight.w500,
//                               color: Colors.red),
//                         ),
//                       ],
//                     )),
//               ),
//               SizedBox(
//                 height: 15,
//               ),
//               TextFormField(
//                 controller: widget.confirmPassword2,
//                 decoration: InputDecoration(
//                   hintText: 'Confirm your Password',
//                   hintStyle: TextStyle(color: Color(0xFFC7C7C7)),
//                   prefixIcon: Padding(
//                     padding: const EdgeInsets.all(12.0),
//                     child: Icon(Icons.lock, color: Colors.grey),
//                   ),
//                   suffixIcon: Padding(
//                     padding: const EdgeInsets.all(12.0),
//                     child: SvgPicture.asset(
//                       'assets/icons/mdi_hide-outline.svg', // Your SVG file path
//                       width: 24, // Adjust width and height as needed
//                       height: 24,
//                     ),
//                   ),
//                   border: OutlineInputBorder(
//                     borderRadius: BorderRadius.all(Radius.circular(15.0)),
//                     borderSide: BorderSide(color: Color(0xFF169C89)),
//                   ),
//                   enabledBorder: OutlineInputBorder(
//                     borderRadius: BorderRadius.all(Radius.circular(15.0)),
//                     borderSide: BorderSide(color: Color(0xFF169C89)),
//                   ),
//                   focusedBorder: OutlineInputBorder(
//                     borderRadius: BorderRadius.all(Radius.circular(15.0)),
//                     borderSide: BorderSide(color: Color(0xFF169C89), width: 2),
//                   ),
//                   filled: true,
//                   fillColor: Color.fromARGB(255, 247, 245, 245),
//                 ),
//                 obscureText: true,
//               ),
//               SizedBox(height: 20),
//               Padding(
//                 padding: EdgeInsets.symmetric(horizontal: 10),
//                 child: const Align(
//                     alignment: Alignment.centerLeft,
//                     child: Row(
//                       children: [
//                         Text(
//                           'Language',
//                           style: TextStyle(
//                               fontSize: 15, fontWeight: FontWeight.w500),
//                         ),
//                         SizedBox(
//                           width: 5,
//                         ),
//                         Text(
//                           '*',
//                           style: TextStyle(
//                               fontSize: 15,
//                               fontWeight: FontWeight.w500,
//                               color: Colors.red),
//                         ),
//                       ],
//                     )),
//               ),
//               SizedBox(
//                 height: 15,
//               ),

//               // Language Dropdown
//               DropdownButtonFormField<String>(
//                 decoration: InputDecoration(
//                   hintText: 'Select your preferred language',
//                   hintStyle: TextStyle(color: Color(0xFFC7C7C7)),
//                   prefixIcon: Padding(
//                     padding: const EdgeInsets.all(12.0),
//                     child: Icon(Icons.language, color: Colors.grey),
//                   ),
//                   border: OutlineInputBorder(
//                     borderRadius: BorderRadius.all(Radius.circular(15.0)),
//                     borderSide: BorderSide(color: Color(0xFF169C89)),
//                   ),
//                   enabledBorder: OutlineInputBorder(
//                     borderRadius: BorderRadius.all(Radius.circular(15.0)),
//                     borderSide: BorderSide(color: Color(0xFF169C89)),
//                   ),
//                   focusedBorder: OutlineInputBorder(
//                     borderRadius: BorderRadius.all(Radius.circular(15.0)),
//                     borderSide: BorderSide(color: Color(0xFF169C89), width: 2),
//                   ),
//                   filled: true,
//                   fillColor: Color.fromARGB(255, 247, 245, 245),
//                 ),
//                 value: _selectedLanguage, // Current value of the dropdown
//                 items: [
//                   DropdownMenuItem(value: 'English', child: Text('English')),
//                   DropdownMenuItem(value: 'Amharic', child: Text('Amharic')),
//                 ],
//                 onChanged: (value) {
//                   setState(() {
//                     _selectedLanguage = value!; // Update selected language
//                   });
//                 },
//               ),
//               SizedBox(height: 20),
//               Padding(
//                 padding: EdgeInsets.symmetric(horizontal: 10),
//                 child: const Align(
//                     alignment: Alignment.centerLeft,
//                     child: Row(
//                       children: [
//                         Text(
//                           'please choose a category',
//                           style: TextStyle(
//                               fontSize: 15, fontWeight: FontWeight.w500),
//                         ),
//                         SizedBox(
//                           width: 5,
//                         ),
//                         Text(
//                           '*',
//                           style: TextStyle(
//                               fontSize: 15,
//                               fontWeight: FontWeight.w500,
//                               color: Colors.red),
//                         ),
//                       ],
//                     )),
//               ),
//               SizedBox(
//                 height: 15,
//               ),

//               // Category Dropdown
//               DropdownButtonFormField<String>(
//                 decoration: InputDecoration(
//                   hintText: 'Category',
//                   hintStyle: TextStyle(color: Color(0xFFC7C7C7)),
//                   border: OutlineInputBorder(
//                     borderRadius: BorderRadius.all(Radius.circular(15.0)),
//                     borderSide: BorderSide(color: Color(0xFF169C89)),
//                   ),
//                   enabledBorder: OutlineInputBorder(
//                     borderRadius: BorderRadius.all(Radius.circular(15.0)),
//                     borderSide: BorderSide(color: Color(0xFF169C89)),
//                   ),
//                   focusedBorder: OutlineInputBorder(
//                     borderRadius: BorderRadius.all(Radius.circular(15.0)),
//                     borderSide: BorderSide(color: Color(0xFF169C89), width: 2),
//                   ),
//                   filled: true,
//                   fillColor: Color.fromARGB(255, 247, 245, 245),
//                 ),
//                 value: _selectedCategory, // Current value of the dropdown
//                 items: [
//                   DropdownMenuItem(
//                       value: 'Victim',
//                       child: Text(
//                         'Victim',
//                       )),
//                   DropdownMenuItem(value: 'General', child: Text('General')),
//                 ],
//                 onChanged: (value) {
//                   setState(() {
//                     _selectedCategory = value!; // Update selected category
//                   });
//                 },
//               ),
//               SizedBox(
//                 height: 20,
//               ),
//             ],
//           ),
//         ),
//       ),
//     );
//   }
// }
