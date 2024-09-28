import 'package:flutter/gestures.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:flutter_svg/flutter_svg.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:safe_haven/features/auth/domain/entities/reset_password_entity.dart';
import 'package:safe_haven/features/auth/presentation/bloc/bloc/auth_bloc_bloc.dart';
import 'package:safe_haven/features/auth/presentation/widgets/custom_button.dart';

class ResetPasswordscreen extends StatefulWidget {
  const ResetPasswordscreen({super.key});
  @override
  State<ResetPasswordscreen> createState() => _ResetPasswordScreen();
}

class _ResetPasswordScreen extends State<ResetPasswordscreen> {
  TextEditingController confirmPassword = TextEditingController();
  TextEditingController password = TextEditingController();

  @override
  Widget build(BuildContext context) {
    return BlocListener<AuthBlocBloc, AuthBlocState>(
      listener: (context, state) {
        if (state is ResetPasswordSuccessState) {
          if (!context.mounted) return;
          ScaffoldMessenger.of(context).showSnackBar(SnackBar(
            content: Text(state.successResetPasswordMessage),
            backgroundColor: Theme.of(context).primaryColor,
          ));
        } else if (state is ResetPasswordErrorState) {
          if (!context.mounted) return;
          ScaffoldMessenger.of(context).showSnackBar(SnackBar(
            content: Text(state.errorResetPasswordMessage),
            backgroundColor: Colors.red,
          ));
        }
      },
      child: Scaffold(
          appBar: AppBar(
            title: const Center(
              child: Text(
                'reset Password',
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
                  CustomresetPasswordForm(
                    ConfirmPassword2: confirmPassword,
                    password2: password,
                  ),
                  BlocBuilder<AuthBlocBloc, AuthBlocState>(
                      builder: (context, state) {
                    return CustomButton(
                        text: 'Update Password',
                        onPressed: () {
                          context.read<AuthBlocBloc>().add(ResetPasswordEvent(
                              resetPasswordEntity: ResetPasswordEntity(
                                  new_password: password.text,
                                  reset_token: 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb2RlIjozMzYxOSwiZW1haWwiOiJMaXlhLmRhbmllbC56ZWxla2VAZ21haWwuY29tIiwiZXhwIjoxNzI3MDk2NzU2fQ.3HcvCYIcX2GiWUhxEGQ5oAZBtogNWmVkZa-E0bUhoaQ')));
                        },
                        bC: 0xFFFFFFFF,
                        col: 0xFF169C89);
                  }),
                  const SizedBox(
                    height: 20,
                  ),
                  Row(
                    mainAxisAlignment: MainAxisAlignment.center,
                    children: [
                      const Text(
                        'Wrong passowrd ',
                        style: TextStyle(color: Colors.grey),
                      ),
                      const SizedBox(
                        height: 20,
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
                                      context, '/resetPassword');
                                })),
                    ],
                  ),
                ]),
              ),
            ),
          )),
    );
  }
}

class CustomresetPasswordForm extends StatefulWidget {
  final TextEditingController password2;
  final TextEditingController ConfirmPassword2;

  const CustomresetPasswordForm({
    super.key,
    required this.password2,
    required this.ConfirmPassword2,
  });

  @override
  State<CustomresetPasswordForm> createState() =>
      _CustomresetPasswordFormState();
}

class _CustomresetPasswordFormState extends State<CustomresetPasswordForm> {
  @override
  Widget build(BuildContext context) {
    return SingleChildScrollView(
        child: Form(
      child: Column(children: [
        // Full Name Field

        Padding(
          padding: EdgeInsets.symmetric(horizontal: 10),
          child: Align(
              alignment: Alignment.centerLeft,
              child: Text(
                'Update Password',
                style: GoogleFonts.robotoSerif(
                    fontSize: 32,
                    fontWeight: FontWeight.bold,
                    color: Color(0xFF169C89)),
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
            hintStyle: TextStyle(color: Color(0xFFC7C7C7)),
            prefixIcon: Padding(
              padding: const EdgeInsets.all(12.0),
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
        Padding(
          padding: EdgeInsets.symmetric(horizontal: 10),
          child: const Align(
            alignment: Alignment.centerLeft,
            child: Text(
              'Confirm password',
              style: TextStyle(fontSize: 15, fontWeight: FontWeight.w500),
            ),
          ),
        ),
        const SizedBox(
          height: 10,
        ),

        TextFormField(
          controller: widget.ConfirmPassword2,
          decoration: const InputDecoration(
            hintText: 'must be the same as the one above',
            hintStyle: TextStyle(color: Color(0xFFC7C7C7)),
            prefixIcon: Padding(
              padding: EdgeInsets.all(12.0),
              child: Icon(Icons.person, color: Colors.grey),
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

        const SizedBox(height: 40),
      ]),
    ));
  }
}
