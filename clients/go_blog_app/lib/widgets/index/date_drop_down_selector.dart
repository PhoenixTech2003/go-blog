import 'package:flutter/material.dart';

class DateDropDownSelector extends StatelessWidget {
  const DateDropDownSelector({super.key});
  @override
  Widget build(BuildContext context) {
    return Row(
      children: [
        Padding(
          padding: const EdgeInsets.only(top: 16.0),
          child: DropdownMenu<String>(
            inputDecorationTheme: const InputDecorationTheme(
              isDense: true,
              contentPadding: EdgeInsets.symmetric(vertical: 6, horizontal: 8),
            ),
            label: Text("Sort by date"),
            dropdownMenuEntries: [],
          ),
        ),
      ],
    );
  }
}
