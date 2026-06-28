import 'package:flutter/material.dart';

class FeaturedBlogCard extends StatelessWidget {
  final String imageUrl;
  final String title;
  final String articleSnippet;
  final String blogTag;
  final String blogDate;
  final bool hasImage;

  const FeaturedBlogCard({
    super.key,
    this.imageUrl = "",
    required this.title,
    required this.articleSnippet,
    required this.blogDate,
    required this.blogTag,
    this.hasImage = true,
  });

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.only(top: 22.0),
      child: Column(
        children: [
          if (hasImage)
            Card(
              shape: RoundedRectangleBorder(
                borderRadius: BorderRadius.circular(22),
              ),
              clipBehavior: Clip.antiAlias,
              child: SizedBox(
                width: 600,
                height: 200,
                child: Image.network(imageUrl, fit: BoxFit.cover),
              ),
            ),
          Padding(
            padding: const EdgeInsets.only(top: 16.0, left: 2),
            child: Row(
              children: [
                Text(
                  blogTag.toUpperCase(),
                  style: Theme.of(context).textTheme.labelMedium?.copyWith(
                    color: Theme.of(context).colorScheme.secondary,
                    fontWeight: FontWeight.bold,
                  ),
                ),
                SizedBox(width: 14),
                Text(
                  blogDate,
                  style: Theme.of(context).textTheme.labelMedium?.copyWith(
                    fontWeight: FontWeight.bold,
                  ),
                ),
              ],
            ),
          ),
          Padding(
            padding: const EdgeInsets.symmetric(vertical: 12),
            child: Row(
              children: [
                Expanded(
                  child: Text(
                    title,
                    style: Theme.of(context).textTheme.headlineSmall,
                  ),
                ),
              ],
            ),
          ),
          Text(articleSnippet),
        ],
      ),
    );
  }
}
