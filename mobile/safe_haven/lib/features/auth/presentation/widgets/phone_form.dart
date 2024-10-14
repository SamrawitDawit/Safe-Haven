import 'package:flutter/gestures.dart';
import 'package:flutter/material.dart';

class CustomPhoneForm2 extends StatefulWidget {
  final TextEditingController fullName2;
  final TextEditingController password2;
  final TextEditingController confirmPassword2;
  final TextEditingController phoneNumber2;

  final ValueChanged<String> onLanguageChanged;
  final ValueChanged<String> onCategoryChanged;
  final VoidCallback ontoggleFirst;
  const CustomPhoneForm2({
    super.key,
    required this.fullName2,
    required this.password2,
    required this.confirmPassword2,
    required this.phoneNumber2,
    required this.onLanguageChanged,
    required this.onCategoryChanged,
    required this.ontoggleFirst,
  });

  @override
  State<CustomPhoneForm2> createState() => _CustomPhoneForm2State();
}

class _CustomPhoneForm2State extends State<CustomPhoneForm2> {
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
              height: 5,
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

            const SizedBox(height: 10),
            Align(
              alignment: Alignment.centerLeft,
              child: RichText(
                  text: TextSpan(
                      text: ' Or Sign up with Email',
                      style: const TextStyle(color: Color(0xFF169C89)),
                      recognizer: TapGestureRecognizer()
                        ..onTap = () {
                          print('pressed got to phone');
                          widget.ontoggleFirst();
                        })),
            ),
            const SizedBox(
              height: 5,
            ),
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
              height: 5,
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
                          print(_phonepasswordVisible);
                          print('eziga');
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
                widget.onLanguageChanged(_selectedCategory);
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
