import 'package:flutter/material.dart';
import 'package:go_blog_app/widgets/index/date_drop_down_selector.dart';
import 'package:go_blog_app/widgets/index/heading.dart';
import 'package:go_blog_app/widgets/index/tab_buttons.dart';

class Body extends StatelessWidget {
  const Body({super.key});
  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: EdgeInsets.all(12.9),
      child: Column(
        children: [Heading(), TabButtons(), DateDropDownSelector()],
      ),
    );
  }
}
