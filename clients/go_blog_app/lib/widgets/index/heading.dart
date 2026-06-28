import 'package:flutter/material.dart';

class Heading extends StatelessWidget {
  const Heading({super.key});
  @override
  Widget build(BuildContext context) {
    return Row(
      children: [
        Text(
          "The Reading List",
          style: Theme.of(context).textTheme.headlineSmall,
        ),
      ],
    );
  }
}
