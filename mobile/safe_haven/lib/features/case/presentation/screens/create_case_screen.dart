import 'dart:io';

import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:image_picker/image_picker.dart';
import 'package:safe_haven/features/auth/presentation/bloc/bloc/auth_bloc_bloc.dart';
import 'package:safe_haven/features/auth/presentation/widgets/custom_button.dart';
import 'package:safe_haven/features/case/domain/entities/case_entity.dart';
import 'package:safe_haven/features/case/presentation/bloc/case_bloc.dart';
import 'package:video_player/video_player.dart';

class CreateCaseScreen extends StatefulWidget {
  const CreateCaseScreen({
    super.key,
  });
  @override
  State<CreateCaseScreen> createState() => _CreateCaseScreenScreen();
}

class _CreateCaseScreenScreen extends State<CreateCaseScreen> {
  TextEditingController title = TextEditingController();
  TextEditingController description = TextEditingController();
  TextEditingController location = TextEditingController();
  String selectedLanguage = 'English';
  String selectedCategory = 'Victim';

  File? _image;

  final ImagePicker _picker = ImagePicker();
  VideoPlayerController? _videoPlayerController;

  Future<void> _initializeVideoPlayer(String path) async {
    if (_videoPlayerController != null) {
      await _videoPlayerController!.dispose();
    }
    // Initialize a new controller
    _videoPlayerController = VideoPlayerController.file(File(path))
      ..initialize().then((_) {
        setState(() {}); // Ensure the UI is updated after video is initialized
        _videoPlayerController?.play(); // Automatically start playing the video
      });
  }

  @override
  void dispose() {
    _videoPlayerController
        ?.dispose(); // Dispose the video player to free resources
    super.dispose();
  }

  Future<void> _pickImage() async {
    final XFile? pickedFile =
        await _picker.pickImage(source: ImageSource.gallery);
    if (pickedFile != null) {
      setState(() {
        _image = File(pickedFile.path);
        // Store the selected image in _image
      });
    }
  }

  File? _video;
  Future<void> _pickVideo() async {
    final XFile? pickedVideo =
        await _picker.pickVideo(source: ImageSource.gallery);
    if (pickedVideo != null) {
      setState(() {
        _video = File(pickedVideo.path);
        _initializeVideoPlayer(pickedVideo.path);
      });
    }
  }

  bool isPhoneCreateCaseScreen = false;

  void toggleForm() {
    print('changed');
    setState(() {
      isPhoneCreateCaseScreen = !isPhoneCreateCaseScreen;
    });
  }

  final _formKey = GlobalKey<FormState>();
  @override
  Widget build(BuildContext context) {
    return BlocListener<AuthBlocBloc, AuthBlocState>(
      listener: (context, state) {
        print("AuthBlocBloc State: $state");
        if (state is LoggedOut) {
          ScaffoldMessenger.of(context).showSnackBar(SnackBar(
            content: const Text('succesfully logged out'),
          ));
        }
      },
      child: BlocListener<CaseBloc, CaseState>(
        listener: (context, state) {
          print(state);
          print("CaseBloc State: $state");

          if (state is CreatedCase) {
            if (!context.mounted) return;
            ScaffoldMessenger.of(context).showSnackBar(SnackBar(
              content: const Text('Case created successfully (in the ui)'),
              backgroundColor: Theme.of(context).primaryColor,
            ));
            // Navigator.pushNamed(context, '/login');
            print('about to reset event');
            context.read<CaseBloc>().add(ResetCaseEvent());
          } else if (state is CreateCaseError) {
            if (!context.mounted) return;
            ScaffoldMessenger.of(context).showSnackBar(SnackBar(
              content: Text(state.errorMessage),
              backgroundColor: Colors.red,
            ));
          }
        },
        child: Scaffold(
          appBar: AppBar(
            title: const Center(
              child: Text(
                'Report an Incident',
                style: TextStyle(color: Color(0xFF169C89), fontSize: 30),
              ),
            ),
            actions: [
              IconButton(
                  onPressed: () {
                    print('pressed');
                    context.read<AuthBlocBloc>().add(LogoutEvent());
                    print(context);
                  },
                  icon: Icon(Icons.logout))
            ],
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
                GestureDetector(
                  onTap: _pickImage,
                  child: Padding(
                    padding: EdgeInsets.fromLTRB(10, 10, 10, 0),
                    child: Container(
                        height: 150,
                        decoration: BoxDecoration(
                          color: Color.fromARGB(255, 231, 226, 226),
                          borderRadius: BorderRadius.all(Radius.circular(10)),
                          image: _image != null
                              ? DecorationImage(
                                  image: FileImage(
                                      _image!), // Show the selected image
                                  fit: BoxFit
                                      .cover, // Ensures image fills the container
                                )
                              : null, // N),
                        ),
                        child: _image == null
                            ? Center(
                                child: Column(
                                mainAxisAlignment: MainAxisAlignment.center,
                                children: [
                                  Icon(Icons.image_outlined),
                                  Text('optional evidentiary image')
                                ],
                              ))
                            : null),
                  ),
                ),
                GestureDetector(
                  onTap: _pickVideo,
                  child: Padding(
                      padding: EdgeInsets.fromLTRB(10, 10, 10, 0),
                      child: Container(
                        height: 150,
                        width: 350,
                        decoration: BoxDecoration(
                          color: Color.fromARGB(255, 231, 226, 226),
                          borderRadius: BorderRadius.all(Radius.circular(30)),
                        ),
                        child: _video == null
                            ? Container(
                                height: 150,
                                decoration: BoxDecoration(
                                  color: Color.fromARGB(255, 231, 226, 226),
                                  borderRadius:
                                      BorderRadius.all(Radius.circular(10)),
                                ),
                                child: Center(
                                  child: Column(
                                    mainAxisAlignment: MainAxisAlignment.center,
                                    children: [
                                      Icon(Icons.image_outlined),
                                      Text('optional evidentiary video'),
                                    ],
                                  ),
                                ),
                              )
                            : Stack(
                                children: [
                                  // Video Player Container
                                  Container(
                                    height:
                                        150, // Match the height of the parent container
                                    width:
                                        350, // Match the width of the parent container
                                    decoration: BoxDecoration(
                                      borderRadius:
                                          BorderRadius.all(Radius.circular(10)),
                                    ),
                                    child: _videoPlayerController != null &&
                                            _videoPlayerController!
                                                .value.isInitialized
                                        ? ClipRRect(
                                            borderRadius: BorderRadius.circular(
                                                10), // Rounded corners
                                            child: FittedBox(
                                              fit: BoxFit
                                                  .cover, // Ensures video fills the container
                                              child: SizedBox(
                                                width: _videoPlayerController!
                                                    .value.size.width,
                                                height: _videoPlayerController!
                                                    .value.size.height,
                                                child: VideoPlayer(
                                                    _videoPlayerController!),
                                              ),
                                            ),
                                          )
                                        : Center(
                                            child:
                                                CircularProgressIndicator(), // Show loading while video is initializing
                                          ),
                                  ),
                                  // Play/Pause Button
                                  Positioned(
                                    bottom: 10,
                                    right: 10,
                                    child: FloatingActionButton(
                                      backgroundColor:
                                          Colors.black.withOpacity(0.5),
                                      onPressed: () {
                                        setState(() {
                                          // Toggle play/pause state
                                          if (_videoPlayerController!
                                              .value.isPlaying) {
                                            _videoPlayerController!.pause();
                                          } else {
                                            _videoPlayerController!.play();
                                          }
                                        });
                                      },
                                      child: Icon(
                                        _videoPlayerController!.value.isPlaying
                                            ? Icons.pause
                                            : Icons.play_arrow,
                                        color: Colors.white,
                                      ),
                                    ),
                                  ),
                                ],
                              ),
                      )),
                ),

                Form(
                  key: _formKey,
                  child: CustomCreateStateForm(
                    title2: title,
                    description2: description,
                    location2: location,
                  ),
                ),
                BlocBuilder<AuthBlocBloc, AuthBlocState>(
                    builder: (context, state) {
                  return CustomButton(
                      text: 'Report',
                      onPressed: () {
                        print('presseds');
                        print(title.text);
                        print(description.text);
                        print('eziiii');
                        print(location.text);
                        print('wbl');
                        print('imgstrdy');

                        if (_formKey.currentState!.validate()) {
                          context.read<CaseBloc>().add(CreateCaseEvent(
                              singleCase: CaseEntity(
                                  id: 'id',
                                  title: title.text,
                                  description: description.text,
                                  location: location.text,
                                  image_url: _image!.path,
                                  video_url: _video!.path)));
                        }
                      },
                      bC: 0xFFFFFFFF,
                      col: 0xFF169C89);
                }),
                const SizedBox(
                  height: 10,
                ),
              ]),
            ),
          ),
        ),
      ),
    );
  }
}

class CustomCreateStateForm extends StatefulWidget {
  final TextEditingController title2;
  final TextEditingController description2;
  final TextEditingController location2;

  const CustomCreateStateForm({
    super.key,
    required this.title2,
    required this.description2,
    required this.location2,
  });

  @override
  State<CustomCreateStateForm> createState() => _CustomCreateStateFormState();
}

class _CustomCreateStateFormState extends State<CustomCreateStateForm> {
  // State variables to track the dropdown values
  String _selectedLanguage = 'English'; // default value for language
  String _selectedCategory = 'Victim'; // default value for category

  String get selectedLanguage => _selectedLanguage;
  String get selectedCategory => _selectedCategory;

  @override
  Widget build(BuildContext context) {
    return SingleChildScrollView(
      child: Padding(
        padding: const EdgeInsets.fromLTRB(10, 20, 10, 30),
        child: Column(
          children: [
            // Full Name Field
            Padding(
              padding: EdgeInsets.symmetric(horizontal: 10),
              child: Align(
                  alignment: Alignment.centerLeft,
                  child: Row(
                    children: [
                      Text(
                        'Title',
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
              controller: widget.title2,
              decoration: const InputDecoration(
                hintText: 'Enter your title',
                hintStyle: TextStyle(color: Color(0xFFC7C7C7)),
                prefixIcon: Padding(
                  padding: EdgeInsets.all(12.0),
                  child: Icon(Icons.title, color: Colors.grey),
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
                        'Description',
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
              controller: widget.description2,
              decoration: const InputDecoration(
                hintText: 'Enter your description',
                hintStyle: TextStyle(color: Color(0xFFC7C7C7)),
                prefixIcon: Padding(
                    padding: EdgeInsets.all(12.0),
                    child: Icon(
                      Icons.note,
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
                        'Location',
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
              controller: widget.location2,
              decoration: InputDecoration(
                hintText: 'Enter your location',
                hintStyle: const TextStyle(color: Color(0xFFC7C7C7)),
                prefixIcon: const Padding(
                  padding: EdgeInsets.all(12.0),
                  child: Icon(Icons.location_on, color: Colors.grey),
                ),
                suffixIcon: Padding(
                  padding: EdgeInsets.fromLTRB(0, 0, 10, 10),
                  child: Container(
                    height: 30,
                    width: 30,
                    // padding: EdgeInsets.fromLTRB(0, 0, 10, bottom),
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
              // to hide description input
            ),
            const SizedBox(height: 10),

            // Confirm Password Field

            const SizedBox(
              height: 5,
            ),
          ],
        ),
      ),
    );
  }
}
