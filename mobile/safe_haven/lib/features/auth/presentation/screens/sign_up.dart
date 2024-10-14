import 'package:flutter/gestures.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:safe_haven/features/auth/domain/entities/sign_up_entity.dart';
import 'package:safe_haven/features/auth/presentation/bloc/bloc/auth_bloc_bloc.dart';
import 'package:safe_haven/features/auth/presentation/widgets/custom_button.dart';
import 'package:safe_haven/features/auth/presentation/widgets/phone_form.dart';

class SignUpscreen extends StatefulWidget {
  const SignUpscreen({super.key});
  @override
  State<SignUpscreen> createState() => _SignUpScreen();
}

class _SignUpScreen extends State<SignUpscreen> {
  TextEditingController fullName = TextEditingController();
  TextEditingController phone = TextEditingController();
  TextEditingController email = TextEditingController();
  TextEditingController password = TextEditingController();
  TextEditingController ConfirmPassword = TextEditingController();
  String selectedLanguage = 'English';
  String selectedCategory = 'Victim';
  bool isPhoneSignUp = false;

  void toggleForm() {
    print('changed');
    setState(() {
      isPhoneSignUp = !isPhoneSignUp;
    });
  }

  final _formKey = GlobalKey<FormState>();
  @override
  Widget build(BuildContext context) {
    return BlocListener<AuthBlocBloc, AuthBlocState>(
      listener: (context, state) {
        if (state is AuthRegisterSuccess) {
          if (!context.mounted) return;
          ScaffoldMessenger.of(context).showSnackBar(SnackBar(
            content: const Text('Account created successfully (in the ui)'),
            backgroundColor: Theme.of(context).primaryColor,
          ));
          Navigator.pushNamed(context, '/login');
        } else if (state is AuthError) {
          if (!context.mounted) return;
          ScaffoldMessenger.of(context).showSnackBar(SnackBar(
            content: Text(
              state.message,
            ),
            backgroundColor: Colors.red,
          ));
        }
      },
      child: Scaffold(
        appBar: AppBar(
          title: const Center(
            child: Text(
              'Register',
              style: TextStyle(color: Color(0xFF169C89), fontSize: 30),
            ),
          ),
        ),
        body: SingleChildScrollView(
          child: Padding(
            padding: const EdgeInsets.fromLTRB(30, 0, 30, 10),
            child: Column(children: [
              // CustomButton2(
              //     text: 'Sign up with Phone Number',
              //     onPressed: () {
              //       // print(fullName);
              //       Navigator.pushNamed(context, '/signupphone');
              //     },
              //     bC: 0xFF169C89,
              //     col: 0xFFFFFFFF),
              Form(
                key: _formKey,
                child: isPhoneSignUp
                    ? CustomPhoneForm2(
                        ontoggleFirst: toggleForm,
                        fullName2: fullName,
                        password2: password,
                        confirmPassword2: ConfirmPassword,
                        phoneNumber2: phone,
                        onLanguageChanged: (language) {
                          setState(() {
                            selectedLanguage = language;
                          });
                        },
                        onCategoryChanged: (category) {
                          setState(() {
                            selectedLanguage = category;
                          });
                        },
                      )
                    : CustomForm(
                        fullName2: fullName,
                        email2: email,
                        password2: password,
                        confirmPassword2: ConfirmPassword,
                        onLanguageChanged: (language) {
                          setState(() {
                            selectedLanguage = language;
                          });
                        },
                        onCategoryChanged: (category) {
                          setState(() {
                            selectedLanguage = category;
                          });
                        },
                        ontoggleFirst: toggleForm,
                      ),
              ),
              BlocBuilder<AuthBlocBloc, AuthBlocState>(
                  builder: (context, state) {
                return CustomButton(
                    text: 'Register',
                    onPressed: () {
                      // print(fullName);

                      if (_formKey.currentState!.validate()) {
                        isPhoneSignUp
                            ? context.read<AuthBlocBloc>().add(RegisterEvent(
                                registrationEntity: SignUpEntity(
                                    language: selectedLanguage,
                                    category: selectedCategory,
                                    fullName: fullName.text,
                                    password: password.text,
                                    phoneNumber: phone.text)))
                            : context.read<AuthBlocBloc>().add(RegisterEvent(
                                registrationEntity: SignUpEntity(
                                    language: selectedLanguage,
                                    category: selectedCategory,
                                    fullName: fullName.text,
                                    password: password.text,
                                    email: email.text)));
                      }
                    },
                    bC: 0xFFFFFFFF,
                    col: 0xFF169C89);
              }),
              const SizedBox(
                height: 10,
              ),
              Row(
                mainAxisAlignment: MainAxisAlignment.center,
                children: [
                  const Text('Have an account? '),
                  RichText(
                      text: TextSpan(
                          text: 'Log In',
                          style: const TextStyle(color: Color(0xFF169C89)),
                          recognizer: TapGestureRecognizer()
                            ..onTap = () {
                              Navigator.pushNamed(context, '/login');
                            })),
                ],
              ),
            ]),
          ),
        ),
      ),
    );
  }
}

class CustomForm extends StatefulWidget {
  final TextEditingController fullName2;
  final TextEditingController password2;
  final TextEditingController confirmPassword2;
  final TextEditingController email2;

  final ValueChanged<String> onLanguageChanged;
  final ValueChanged<String> onCategoryChanged;
  final VoidCallback ontoggleFirst;
  const CustomForm({
    super.key,
    required this.fullName2,
    required this.password2,
    required this.confirmPassword2,
    required this.email2,
    required this.onLanguageChanged,
    required this.onCategoryChanged,
    required this.ontoggleFirst,
  });

  @override
  State<CustomForm> createState() => _CustomFormState();
}

class _CustomFormState extends State<CustomForm> {
  // State variables to track the dropdown values
  String _selectedLanguage = 'English'; // default value for language
  String _selectedCategory = 'Victim'; // default value for category

  String get selectedLanguage => _selectedLanguage;
  String get selectedCategory => _selectedCategory;

  bool _passwordVisible = false;
  bool _confirmPasswordVisible = false;

  @override
  Widget build(BuildContext context) {
    return SingleChildScrollView(
      child: Padding(
        padding: const EdgeInsets.fromLTRB(10, 50, 10, 30),
        child: Column(
          children: [
            // Full Name Field
            const Padding(
              padding: EdgeInsets.symmetric(horizontal: 10),
              child: Align(
                  alignment: Alignment.centerLeft,
                  child: Row(
                    children: [
                      Text(
                        'Full Name',
                        style: TextStyle(
                            fontSize: 15, fontWeight: FontWeight.w500),
                      ),
                      SizedBox(
                        width: 5,
                      ),
                      Text(
                        '*',
                        style: TextStyle(
                            fontSize: 15,
                            fontWeight: FontWeight.w500,
                            color: Colors.red),
                      ),
                    ],
                  )),
            ),
            const SizedBox(
              height: 5,
            ),
            TextFormField(
              validator: (value) {
                if (value == null || value.isEmpty) {
                  return 'Full Name cannot be empty';
                }
                return null;
              },
              controller: widget.fullName2,
              decoration: const InputDecoration(
                hintText: 'Enter your name',
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

            const SizedBox(height: 10),
            Align(
              alignment: Alignment.centerLeft,
              child: RichText(
                  text: TextSpan(
                      text: ' Or Sign up with phone',
                      style: const TextStyle(color: Color(0xFF169C89)),
                      recognizer: TapGestureRecognizer()
                        ..onTap = () {
                          print('pressed got to phone');
                          widget.ontoggleFirst();
                        })),
            ),
            SizedBox(
              height: 5,
            ),
            const Padding(
              padding: EdgeInsets.symmetric(horizontal: 10),
              child: Align(
                  alignment: Alignment.centerLeft,
                  child: Row(
                    children: [
                      Text(
                        'Email address',
                        style: TextStyle(
                            fontSize: 15, fontWeight: FontWeight.w500),
                      ),
                      SizedBox(
                        width: 5,
                      ),
                      Text(
                        '*',
                        style: TextStyle(
                            fontSize: 15,
                            fontWeight: FontWeight.w500,
                            color: Colors.red),
                      ),
                    ],
                  )),
            ),
            const SizedBox(
              height: 5,
            ),

            TextFormField(
              validator: (value) {
                if (value == null || value.isEmpty) {
                  return 'Email cannot be empty';
                }
                return null;
              },
              controller: widget.email2,
              decoration: const InputDecoration(
                hintText: 'Enter your email',
                hintStyle: TextStyle(color: Color(0xFFC7C7C7)),
                prefixIcon: Padding(
                    padding: EdgeInsets.all(12.0),
                    child: Icon(
                      Icons.email,
                      color: Colors.grey,
                    )),
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
            const SizedBox(height: 10),
            const Padding(
              padding: EdgeInsets.symmetric(horizontal: 10),
              child: Align(
                  alignment: Alignment.centerLeft,
                  child: Row(
                    children: [
                      Text(
                        'Password',
                        style: TextStyle(
                            fontSize: 15, fontWeight: FontWeight.w500),
                      ),
                      SizedBox(
                        width: 5,
                      ),
                      Text(
                        '*',
                        style: TextStyle(
                            fontSize: 15,
                            fontWeight: FontWeight.w500,
                            color: Colors.red),
                      ),
                    ],
                  )),
            ),
            const SizedBox(
              height: 5,
            ),

            // Password Field
            TextFormField(
              validator: (value) {
                if (value == null || value.isEmpty) {
                  return 'Password cannot be empty';
                }
                // Regular expression to check password criteria
                final hasUppercase = RegExp(r'[A-Z]');
                final hasDigits = RegExp(r'\d');
                final hasSpecialCharacters = RegExp(r'[!@#$%^&*(),.?":{}|<>]');

                if (!hasUppercase.hasMatch(value)) {
                  return 'Password must contain at least one uppercase letter';
                }
                if (!hasDigits.hasMatch(value)) {
                  return 'Password must contain at least one number';
                }
                if (!hasSpecialCharacters.hasMatch(value)) {
                  return 'Password must contain at least one special character';
                }
                if (value.length < 8) {
                  return 'Password must be at least 8 characters long';
                }

                return null;
              },

              obscureText: !_passwordVisible,
              controller: widget.password2,
              decoration: InputDecoration(
                hintText: 'Enter your password',
                hintStyle: const TextStyle(color: Color(0xFFC7C7C7)),
                prefixIcon: const Padding(
                  padding: EdgeInsets.all(12.0),
                  child: Icon(Icons.lock, color: Colors.grey),
                ),
                suffixIcon: Padding(
                  padding: EdgeInsets.fromLTRB(0, 0, 10, 10),
                  child: Container(
                    height: 30,
                    width: 30,
                    // padding: EdgeInsets.fromLTRB(0, 0, 10, bottom),
                    child: IconButton(
                      icon: Icon(
                        _passwordVisible
                            ? Icons.visibility
                            : Icons.visibility_off,
                        color: Colors.grey,
                      ),
                      onPressed: () {
                        setState(() {
                          _passwordVisible =
                              !_passwordVisible; // Toggle visibility
                        });
                      },
                    ),
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
              // to hide password input
            ),
            const SizedBox(height: 10),

            // Confirm Password Field
            const Padding(
              padding: EdgeInsets.symmetric(horizontal: 10),
              child: Align(
                  alignment: Alignment.centerLeft,
                  child: Row(
                    children: [
                      Text(
                        'Confirm Password',
                        style: TextStyle(
                            fontSize: 15, fontWeight: FontWeight.w500),
                      ),
                      SizedBox(
                        width: 5,
                      ),
                      Text(
                        '*',
                        style: TextStyle(
                            fontSize: 15,
                            fontWeight: FontWeight.w500,
                            color: Colors.red),
                      ),
                    ],
                  )),
            ),
            const SizedBox(
              height: 5,
            ),
            TextFormField(
              validator: (value) {
                if (value == null || value.isEmpty) {
                  return 'Please confirm your password';
                } else if (value != widget.password2.text) {
                  return 'Passwords do not match';
                }
                return null;
              },
              obscureText: !_confirmPasswordVisible,
              controller: widget.confirmPassword2,
              decoration: InputDecoration(
                hintText: 'Confirm your Password',
                hintStyle: const TextStyle(color: Color(0xFFC7C7C7)),
                prefixIcon: const Padding(
                  padding: EdgeInsets.all(12.0),
                  child: Icon(Icons.lock, color: Colors.grey),
                ),
                suffixIcon: Padding(
                  padding: EdgeInsets.fromLTRB(0, 0, 10, 10),
                  child: Container(
                    height: 30,
                    width: 30,
                    child: IconButton(
                      icon: Icon(
                        _confirmPasswordVisible
                            ? Icons.visibility
                            : Icons.visibility_off,
                        color: Colors.grey,
                      ),
                      onPressed: () {
                        setState(() {
                          _confirmPasswordVisible =
                              !_confirmPasswordVisible; // Toggle visibility
                        });
                      },
                    ),
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
            const SizedBox(height: 10),
            const Padding(
              padding: EdgeInsets.symmetric(horizontal: 10),
              child: Align(
                  alignment: Alignment.centerLeft,
                  child: Row(
                    children: [
                      Text(
                        'Language',
                        style: TextStyle(
                            fontSize: 15, fontWeight: FontWeight.w500),
                      ),
                      SizedBox(
                        width: 5,
                      ),
                      Text(
                        '*',
                        style: TextStyle(
                            fontSize: 15,
                            fontWeight: FontWeight.w500,
                            color: Colors.red),
                      ),
                    ],
                  )),
            ),
            const SizedBox(
              height: 5,
            ),

            // Language Dropdown
            DropdownButtonFormField<String>(
              decoration: const InputDecoration(
                hintText: 'Select your preferred language',
                hintStyle: TextStyle(color: Color(0xFFC7C7C7)),
                prefixIcon: Padding(
                  padding: EdgeInsets.all(12.0),
                  child: Icon(Icons.language, color: Colors.grey),
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
              value: _selectedLanguage, // Current value of the dropdown
              items: const [
                DropdownMenuItem(value: 'English', child: Text('English')),
                DropdownMenuItem(value: 'Amharic', child: Text('Amharic')),
              ],
              onChanged: (value) {
                setState(() {
                  _selectedLanguage = value!; // Update selected language
                });
                widget.onLanguageChanged(_selectedLanguage);
              },
            ),
            const SizedBox(height: 10),
            const Padding(
              padding: EdgeInsets.symmetric(horizontal: 10),
              child: Align(
                  alignment: Alignment.centerLeft,
                  child: Row(
                    children: [
                      Text(
                        'please choose a category',
                        style: TextStyle(
                            fontSize: 15, fontWeight: FontWeight.w500),
                      ),
                      SizedBox(
                        width: 5,
                      ),
                      Text(
                        '*',
                        style: TextStyle(
                            fontSize: 15,
                            fontWeight: FontWeight.w500,
                            color: Colors.red),
                      ),
                    ],
                  )),
            ),
            const SizedBox(
              height: 5,
            ),

            // Category Dropdown
            DropdownButtonFormField<String>(
              decoration: const InputDecoration(
                hintText: 'Category',
                hintStyle: TextStyle(color: Color(0xFFC7C7C7)),
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
              value: _selectedCategory, // Current value of the dropdown
              items: const [
                DropdownMenuItem(
                    value: 'Victim',
                    child: Text(
                      'Victim',
                    )),
                DropdownMenuItem(value: 'General', child: Text('General')),
              ],
              onChanged: (value) {
                setState(() {
                  _selectedCategory = value!; // Update selected category
                });
                widget.onCategoryChanged(_selectedCategory);
              },
            ),
          ],
        ),
      ),
    );
  }
}
