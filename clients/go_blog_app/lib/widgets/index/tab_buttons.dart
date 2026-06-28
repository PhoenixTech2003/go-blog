import 'package:flutter/material.dart';

class TabButtons extends StatelessWidget {
  const TabButtons({super.key});
  @override
  Widget build(BuildContext context) {
    const tabItems = ["Design", "Tech", "Life"];
    return Padding(
      padding: const EdgeInsets.only(top: 12),
      child: Row(
        children: [
          TextButton(
            onPressed: () {},
            style: TextButton.styleFrom(
              backgroundColor: Theme.of(context).primaryColor,
              foregroundColor: Theme.of(context).canvasColor,
            ),
            child: const Text("All"),
          ),
          for (var tabItem in tabItems)
            TextButton(
              onPressed: () {},
              style: TextButton.styleFrom(
                foregroundColor: Theme.of(context).primaryColor,
              ),
              child: Text(tabItem),
            ),
        ],
      ),
    );
  }
}
