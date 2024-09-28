import 'package:flutter/gestures.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:flutter_svg/flutter_svg.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:safe_haven/features/auth/domain/entities/log_in_entity.dart';
import 'package:safe_haven/features/auth/presentation/bloc/bloc/auth_bloc_bloc.dart';
import 'package:safe_haven/features/auth/presentation/widgets/custom_button.dart';

class LogInscreen extends StatefulWidget {
  const LogInscreen({super.key});
  @override
  State<LogInscreen> createState() => _LogInScreen();
}

class _LogInScreen extends State<LogInscreen> {
  TextEditingController email = TextEditingController();
  TextEditingController password = TextEditingController();

  @override
  Widget build(BuildContext context) {
    return BlocListener<AuthBlocBloc, AuthBlocState>(
      listener: (context, state) {
        if (state is LogInSuccess) {
          if (!context.mounted) return;
          ScaffoldMessenger.of(context).showSnackBar(SnackBar(
            content: const Text('Successfully logged in (in the ui)'),
            backgroundColor: Theme.of(context).primaryColor,
          ));
        } else if (state is LoggInError) {
          if (!context.mounted) return;
          ScaffoldMessenger.of(context).showSnackBar(SnackBar(
            content: Text(state.logInErrorMessage),
            backgroundColor: Colors.red,
          ));
        }
      },
      child: Scaffold(
          appBar: AppBar(
            leading: IconButton(
                onPressed: () {
                  Navigator.pop(context);
                },
                icon: const Icon(Icons.arrow_back)),
            title: const Center(
              child: Text(
                'Login',
                style: TextStyle(color: Color(0xFF169C89)),
              ),
            ),
          ),
          body: SingleChildScrollView(
            child: Padding(
              padding: const EdgeInsets.fromLTRB(30, 30, 30, 10),
              child: Padding(
                padding: const EdgeInsets.symmetric(vertical: 70),
                child: Column(children: [
                  CustomLoginForm(
                    email2: email,
                    password2: password,
                  ),
                  BlocBuilder<AuthBlocBloc, AuthBlocState>(
                      builder: (context, state) {
                    return CustomButton(
                        text: 'Log In',
                        onPressed: () {
                          context.read<AuthBlocBloc>().add(LoginEvent(
                              loginEntity: LogInEntity(
                                  userType: 'normal',
                                  password: password.text,
                                  email: email.text)));
                        },
                        bC: 0xFFFFFFFF,
                        col: 0xFF169C89);
                  }),
                  const SizedBox(
                    height: 20,
                  ),
                  Row(
                    mainAxisAlignment: MainAxisAlignment.spaceBetween,
                    children: [
                      Row(
                        children: [
                          const Text('Don\'t have an account?'),
                          const SizedBox(
                            width: 3,
                          ),
                          RichText(
                              text: TextSpan(
                                  text: 'Sign up',
                                  style: const TextStyle(
                                    color: Color(0xFF169C89),
                                  ),
                                  recognizer: TapGestureRecognizer()
                                    ..onTap = () {
                                      Navigator.pushNamed(context, '/signup');
                                    }))
                        ],
                      ),
                      RichText(
                          text: TextSpan(
                              text: 'Forgot password?',
                              style: const TextStyle(
                                color: Color(0xFF169C89),
                              ),
                              recognizer: TapGestureRecognizer()
                                ..onTap = () {
                                  Navigator.pushNamed(
                                      context, '/forgotpassword');
                                })),
                    ],
                  ),
                  const SizedBox(
                    height: 10,
                  ),
                  const Text(
                    'Or connect with google',
                    style: TextStyle(
                      color: Colors.grey,
                    ),
                  ),
                  const SizedBox(
                    height: 3,
                  ),
                  CustomButton3(
                      widget: SvgPicture.asset('assets/icons/google_logo.svg'),
                      onPressed: () {
                        context
                            .read<AuthBlocBloc>()
                            .add(const GoogleSignInEvent());
                      },
                      bC: 0xFF169C89,
                      col: 0xFFFFFFFF)
                ]),
              ),
            ),
          )),
    );
  }
}

class CustomLoginForm extends StatefulWidget {
  final TextEditingController password2;
  final TextEditingController email2;

  const CustomLoginForm({
    super.key,
    required this.password2,
    required this.email2,
  });

  @override
  State<CustomLoginForm> createState() => _CustomLoginFormState();
}

class _CustomLoginFormState extends State<CustomLoginForm> {
  @override
  Widget build(BuildContext context) {
    return SingleChildScrollView(
        child: Form(
      child: Column(children: [
        // Full Name Field

        Padding(
          padding: const EdgeInsets.symmetric(horizontal: 10),
          child: Align(
              alignment: Alignment.centerLeft,
              child: Text(
                'Log In',
                style: GoogleFonts.robotoSerif(
                    fontSize: 32,
                    fontWeight: FontWeight.bold,
                    color: const Color(0xFF169C89)),
              )),
        ),
        const SizedBox(
          height: 30,
        ),

        const Padding(
          padding: EdgeInsets.symmetric(horizontal: 10),
          child: Align(
            alignment: Alignment.centerLeft,
            child: Text(
              'Email address',
              style: TextStyle(fontSize: 15, fontWeight: FontWeight.w500),
            ),
          ),
        ),
        const SizedBox(
          height: 10,
        ),

        TextFormField(
          controller: widget.email2,
          decoration: const InputDecoration(
            hintText: 'Email',
            hintStyle: TextStyle(color: Color(0xFFC7C7C7)),
            prefixIcon: Padding(
              padding: EdgeInsets.all(12.0),
              child: Icon(Icons.email, color: Colors.grey),
            ),
            border: OutlineInputBorder(
              borderRadius: BorderRadius.all(Radius.circular(15.0)),
              borderSide: BorderSide(color: Color(0xFF169C89)),
            ),
            enabledBorder: OutlineInputBorder(
              borderRadius: BorderRadius.all(Radius.circular(15.0)),
              borderSide: BorderSide(color: Color(0xFF169C89)),
            ),
            focusedBorder: OutlineInputBorder(
              borderRadius: BorderRadius.all(Radius.circular(15.0)),
              borderSide: BorderSide(color: Color(0xFF169C89), width: 2),
            ),
            filled: true,
            fillColor: Color.fromARGB(255, 247, 245, 245),
          ),
        ),
        const SizedBox(height: 20),
        const Padding(
          padding: EdgeInsets.symmetric(horizontal: 10),
          child: Align(
            alignment: Alignment.centerLeft,
            child: Text(
              'Password',
              style: TextStyle(fontSize: 15, fontWeight: FontWeight.w500),
            ),
          ),
        ),
        const SizedBox(
          height: 10,
        ),
        TextFormField(
          controller: widget.password2,
          decoration: InputDecoration(
            hintText: 'password',
            hintStyle: const TextStyle(color: Color(0xFFC7C7C7)),
            prefixIcon: const Padding(
              padding: EdgeInsets.all(12.0),
              child: Icon(Icons.lock, color: Colors.grey),
            ),
            suffixIcon: Padding(
              padding: const EdgeInsets.all(12.0),
              child: SvgPicture.asset(
                'assets/icons/mdi_hide-outline.svg', // Your SVG file path
                width: 24, // Adjust width and height as needed
                height: 24,
              ),
            ),
            border: const OutlineInputBorder(
              borderRadius: BorderRadius.all(Radius.circular(15.0)),
              borderSide: BorderSide(color: Color(0xFF169C89)),
            ),
            enabledBorder: const OutlineInputBorder(
              borderRadius: BorderRadius.all(Radius.circular(15.0)),
              borderSide: BorderSide(color: Color(0xFF169C89)),
            ),
            focusedBorder: const OutlineInputBorder(
              borderRadius: BorderRadius.all(Radius.circular(15.0)),
              borderSide: BorderSide(color: Color(0xFF169C89), width: 2),
            ),
            filled: true,
            fillColor: const Color.fromARGB(255, 247, 245, 245),
          ),
        ),

        const SizedBox(height: 40),
      ]),
    ));
  }
}
