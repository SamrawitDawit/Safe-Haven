import 'dart:ui';

import 'package:flutter/material.dart';

class CustomButton extends StatelessWidget {
  final String text;
  final VoidCallback onPressed;
  final bC;
  final col;

  const CustomButton(
      {super.key,
      required this.text,
      required this.onPressed,
      required this.bC,
      required this.col});

  @override
  Widget build(BuildContext context) {
    return Container(
        height: 50,
        width: 500,
        decoration: BoxDecoration(
            color: Color(col),
            border: Border.all(color: Color(bC)),
            borderRadius: BorderRadius.circular(8)),
        child: Material(
          color: Colors.transparent,
          child: InkWell(
            onTap: onPressed,
            child: Center(
                child: Text(
              text,
              style: TextStyle(fontSize: 20, color: Color(bC)),
            )),
          ),
        ));
  }
}

class CustomButton2 extends StatelessWidget {
  final String text;
  final VoidCallback onPressed;
  final int bC; // Color for the border
  final int col; // Background color

  const CustomButton2({
    super.key,
    required this.text,
    required this.onPressed,
    required this.bC,
    required this.col,
  });

  @override
  Widget build(BuildContext context) {
    return Material(
      color: Colors.transparent,
      child: InkWell(
        onTap: onPressed,
        child: Container(
          padding: const EdgeInsets.symmetric(
              horizontal: 12, vertical: 8), // Adjusted padding
          decoration: BoxDecoration(
            color: Color(col),
            border: Border.all(color: Color(bC)),
            borderRadius: BorderRadius.circular(8),
          ),
          child: FittedBox(
            // Ensures button resizes tightly around the text
            child: Text(
              text,
              style: TextStyle(fontSize: 15, color: Color(bC)),
            ),
          ),
        ),
      ),
    );
  }
}

class CustomButton3 extends StatelessWidget {
  final Widget widget;
  final VoidCallback onPressed;
  final int bC; // Color for the border
  final int col; // Background color

  const CustomButton3({
    super.key,
    required this.widget,
    required this.onPressed,
    required this.bC,
    required this.col,
  });

  @override
  Widget build(BuildContext context) {
    return Material(
      color: Colors.transparent,
      child: InkWell(
        onTap: onPressed,
        child: Container(
          padding: const EdgeInsets.symmetric(
              horizontal: 8, vertical: 3), // Adjusted padding
          decoration: BoxDecoration(
            color: Color(col),
            border: Border.all(color: Color(bC)),
            borderRadius: BorderRadius.circular(8),
          ),
          child: FittedBox(
              // Ensures button resizes tightly around the text
              child: IconButton(onPressed: () {}, icon: widget)),
        ),
      ),
    );
  }
}
