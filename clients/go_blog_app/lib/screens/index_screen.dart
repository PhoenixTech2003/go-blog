import 'package:flutter/material.dart';
import 'package:go_blog_app/widgets/index/body.dart';

class IndexScreen extends StatelessWidget {
  const IndexScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text("Minimalist"),

        actions: [
          Padding(
            padding: const EdgeInsets.only(right: 8),
            child: TextButton.icon(
              onPressed: () {},
              style: TextButton.styleFrom(
                backgroundColor: Theme.of(context).primaryColor,
                foregroundColor: Theme.of(context).cardColor,
                shape: RoundedRectangleBorder(
                  borderRadius: BorderRadius.circular(0),
                ),
              ),
              label: const Text("New Post"),
              icon: const Icon(Icons.add),
            ),
          ),
          IconButton(
            onPressed: () {},
            style: IconButton.styleFrom(
              side: BorderSide(
                color: Theme.of(context).primaryColor,
                width: 2.0,
              ),
            ),
            icon: const Icon(Icons.person_3_outlined),
          ),
        ],
      ),
      body: Body(),
    );
  }
}
