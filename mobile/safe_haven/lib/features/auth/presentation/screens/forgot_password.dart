import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:safe_haven/features/auth/presentation/bloc/bloc/auth_bloc_bloc.dart';
import 'package:safe_haven/features/auth/presentation/widgets/custom_button.dart';

class ForgotPasswordscreen extends StatefulWidget {
  const ForgotPasswordscreen({super.key});
  @override
  State<ForgotPasswordscreen> createState() => _ForgotPasswordScreen();
}

class _ForgotPasswordScreen extends State<ForgotPasswordscreen> {
  TextEditingController email = TextEditingController();
  TextEditingController password = TextEditingController();

  @override
  Widget build(BuildContext context) {
    return BlocListener<AuthBlocBloc, AuthBlocState>(
      listener: (context, state) {
        if (state is ForgotPasswordSuccess) {
          if (!context.mounted) return;
          ScaffoldMessenger.of(context).showSnackBar(SnackBar(
            content: const Text('Successfully sent reset email in (in the ui)'),
            backgroundColor: Theme.of(context).primaryColor,
          ));
        } else if (state is ForgotPasswordError) {
          if (!context.mounted) return;
          ScaffoldMessenger.of(context).showSnackBar(SnackBar(
            content: Text(state.forgotPasswordErrorMessage),
            backgroundColor: Colors.red,
          ));
        }
      },
      child: Scaffold(
          appBar: AppBar(
            leading: IconButton(
                onPressed: () {
                  print('pressed');
                  Navigator.pop(context);
                },
                icon: const Icon(Icons.arrow_back_outlined)),
          ),
          body: SingleChildScrollView(
            child: Padding(
              padding: const EdgeInsets.fromLTRB(30, 30, 30, 10),
              child: Padding(
                padding: const EdgeInsets.symmetric(vertical: 70),
                child: Column(children: [
                  CustomforgotPasswordForm(
                    email2: email,
                  ),
                  BlocBuilder<AuthBlocBloc, AuthBlocState>(
                      builder: (context, state) {
                    return CustomButton(
                        text: 'Send email',
                        onPressed: () {
                          context
                              .read<AuthBlocBloc>()
                              .add(ForgotPasswordEvent(resetEmail: email.text));
                        },
                        bC: 0xFFFFFFFF,
                        col: 0xFF169C89);
                  }),
                  const SizedBox(
                    height: 20,
                  ),
                ]),
              ),
            ),
          )),
    );
  }
}

class CustomforgotPasswordForm extends StatefulWidget {
  final TextEditingController email2;

  const CustomforgotPasswordForm({
    super.key,
    required this.email2,
  });

  @override
  State<CustomforgotPasswordForm> createState() =>
      _CustomforgotPasswordFormState();
}

class _CustomforgotPasswordFormState extends State<CustomforgotPasswordForm> {
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
                'Forgot Password?',
                style: GoogleFonts.robotoSerif(
                    fontSize: 32,
                    fontWeight: FontWeight.bold,
                    color: const Color(0xFF169C89)),
              )),
        ),
        const SizedBox(
          height: 15,
        ),
        const Padding(
          padding: EdgeInsets.symmetric(horizontal: 10),
          child: Align(
            alignment: Alignment.centerLeft,
            child: Text(
              'please enter your email to reset the password',
              style: TextStyle(color: Colors.grey),
            ),
          ),
        ),

        const SizedBox(
          height: 15,
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
            hintText: 'Email ',
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

        const SizedBox(height: 40),
      ]),
    ));
  }
}
