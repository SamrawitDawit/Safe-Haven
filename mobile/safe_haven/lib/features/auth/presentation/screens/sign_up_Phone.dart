import 'package:flutter/gestures.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:safe_haven/features/auth/domain/entities/sign_up_entity.dart';
import 'package:safe_haven/features/auth/presentation/bloc/bloc/auth_bloc_bloc.dart';
import 'package:safe_haven/features/auth/presentation/widgets/custom_button.dart';

class SignUpPhonescreen extends StatefulWidget {
  const SignUpPhonescreen({super.key});
  @override
  State<SignUpPhonescreen> createState() => _SignUpPhoneScreen();
}

class _SignUpPhoneScreen extends State<SignUpPhonescreen> {
  TextEditingController fullName = TextEditingController();
  TextEditingController phoneNumber = TextEditingController();
  TextEditingController password = TextEditingController();
  TextEditingController ConfirmPassword = TextEditingController();
  TextEditingController language = TextEditingController();
  TextEditingController category = TextEditingController();

  String selectedLanguage = 'English';
  String selectedCategory = 'General';

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
            content: Text(state.message),
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
            padding: const EdgeInsets.fromLTRB(30, 30, 30, 10),
            child: Column(children: [
              CustomButton2(
                  text: 'Sign up with email',
                  onPressed: () {
                    // print(fullName);
                    Navigator.pushNamed(context, '/signup');
                  },
                  bC: 0xFF169C89,
                  col: 0xFFFFFFFF),
              Form(
                key: _formKey,
                child: CustomPhoneForm(
                  fullName2: fullName,
                  phoneNumber2: phoneNumber,
                  password2: password,
                  confirmPassword2: ConfirmPassword,
                  onLanguageChanged: (language) {
                    selectedLanguage = language;
                  },
                  onCategoryChanged: (category) {
                    selectedCategory = category;
                  },
                ),
              ),
              BlocBuilder<AuthBlocBloc, AuthBlocState>(
                  builder: (context, state) {
                return CustomButton(
                    text: 'Register',
                    onPressed: () {
                      // print(fullName);
                      if (_formKey.currentState!.validate()) {
                        context.read<AuthBlocBloc>().add(RegisterEvent(
                            registrationEntity: SignUpEntity(
                                language: 'language',
                                category: 'category',
                                fullName: fullName.text,
                                password: password.text,
                                phoneNumber: phoneNumber.text)));
                      }
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

class CustomPhoneForm extends StatefulWidget {
  final TextEditingController fullName2;
  final TextEditingController password2;
  final TextEditingController confirmPassword2;
  final TextEditingController phoneNumber2;

  final ValueChanged<String> onLanguageChanged;
  final ValueChanged<String> onCategoryChanged;
  const CustomPhoneForm({
    super.key,
    required this.fullName2,
    required this.password2,
    required this.confirmPassword2,
    required this.phoneNumber2,
    required this.onLanguageChanged,
    required this.onCategoryChanged,
  });

  @override
  State<CustomPhoneForm> createState() => _CustomPhoneFormState();
}

class _CustomPhoneFormState extends State<CustomPhoneForm> {
  // State variables to track the dropdown values
  String _selectedLanguage = 'English'; // default value for language
  String _selectedCategory = 'Victim'; // default value for category

  String get selectedLanguage => _selectedLanguage;
  String get selectedCategory => _selectedCategory;

  bool _phonepasswordVisible = false;
  bool _phoneconfirmPasswordVisible = false;

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
              height: 15,
            ),
            TextFormField(
              validator: (value) {
                if (value == null || value.isEmpty) {
                  return 'Full name can not be empty';
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

            const SizedBox(height: 20),

            const Padding(
              padding: EdgeInsets.symmetric(horizontal: 10),
              child: Align(
                  alignment: Alignment.centerLeft,
                  child: Row(
                    children: [
                      Text(
                        'Phone Number ',
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
              height: 15,
            ),
            TextFormField(
              validator: (value) {
                if (value == null || value.isEmpty) {
                  return 'Phone Number can not be empty';
                }
                return null;
              },
              controller: widget.phoneNumber2,
              decoration: const InputDecoration(
                hintText: 'Enter your Phone Number',
                hintStyle: TextStyle(color: Color(0xFFC7C7C7)),
                prefixIcon: Padding(
                    padding: EdgeInsets.all(12.0),
                    child: Icon(
                      Icons.phone,
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
            const SizedBox(height: 20),
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
              height: 15,
            ),

            // Password Field
            TextFormField(
              validator: (value) {
                if (value == null || value.isEmpty) {
                  return 'Password can not be empty';
                }
                return null;
              },
              obscureText: !_phonepasswordVisible,
              controller: widget.password2,
              decoration: InputDecoration(
                hintText: 'Enter your password',
                hintStyle: const TextStyle(color: Color(0xFFC7C7C7)),
                prefixIcon: const Padding(
                  padding: EdgeInsets.all(12.0),
                  child: Icon(Icons.lock, color: Colors.grey),
                ),
                suffixIcon: Padding(
                  padding: const EdgeInsets.all(12.0),
                  child: Container(
                    height: 30,
                    width: 30,
                    // padding: EdgeInsets.fromLTRB(0, 0, 10, bottom),
                    child: IconButton(
                      icon: Icon(
                        _phonepasswordVisible
                            ? Icons.visibility
                            : Icons.visibility_off,
                        color: Colors.grey,
                      ),
                      onPressed: () {
                        setState(() {
                          print('password eziga');
                          print(_phonepasswordVisible);
                          _phonepasswordVisible =
                              !_phonepasswordVisible; // Toggle visibility
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
              ), // to hide password input
            ),
            const SizedBox(height: 20),

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
              height: 15,
            ),
            TextFormField(
              validator: (value) {
                if (value == null || value.isEmpty) {
                  return 'Please confirm your password';
                }
                return null;
              },
              obscureText: !_phoneconfirmPasswordVisible,
              controller: widget.confirmPassword2,
              decoration: InputDecoration(
                hintText: 'Confirm your Password',
                hintStyle: const TextStyle(color: Color(0xFFC7C7C7)),
                prefixIcon: const Padding(
                  padding: EdgeInsets.all(12.0),
                  child: Icon(Icons.lock, color: Colors.grey),
                ),
                suffixIcon: Padding(
                  padding: const EdgeInsets.all(12.0),
                  child: Container(
                    padding: EdgeInsets.fromLTRB(0, 0, 10, 10),
                    height: 30,
                    width: 30,
                    child: IconButton(
                      icon: Icon(
                        _phoneconfirmPasswordVisible
                            ? Icons.visibility
                            : Icons.visibility_off,
                        color: Colors.grey,
                      ),
                      onPressed: () {
                        setState(() {
                          print('eziga');
                          _phoneconfirmPasswordVisible =
                              !_phoneconfirmPasswordVisible; // Toggle visibility
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
            const SizedBox(height: 20),
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
              height: 15,
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
                widget.onLanguageChanged(_selectedCategory);
              },
            ),
            const SizedBox(height: 20),
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
              height: 15,
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
            const SizedBox(
              height: 20,
            ),
          ],
        ),
      ),
    );
  }
}
