import 'package:flutter/material.dart';
import 'package:go_blog_app/widgets/index/date_drop_down_selector.dart';
import 'package:go_blog_app/widgets/index/featured_blog_card.dart';
import 'package:go_blog_app/widgets/index/heading.dart';
import 'package:go_blog_app/widgets/index/tab_buttons.dart';

class Body extends StatelessWidget {
  const Body({super.key});
  @override
  Widget build(BuildContext context) {
    return SafeArea(
      bottom: true,
      child: SingleChildScrollView(
        child: Padding(
          padding: EdgeInsets.all(12.9),
          child: Column(
            children: [
              Heading(),
              TabButtons(),
              DateDropDownSelector(),
              FeaturedBlogCard(
                imageUrl:
                    "https://lh3.googleusercontent.com/aida-public/AB6AXuDpm-bgBiNwGjJbfIe0oOfnap5Tyw96rl0CRwkb3cPSa2aFnltbWUQlcWqNYDkV9QerZtdndECkLLcuUzo_yHBFN-5uAJF5vbEg_bOetCv7glOJ-fSRYH0i1Qy-pbo2_8L90K-SXHvBk4UnAQH2NED2StAHYPOcZt4DsQiDQi4-PTGFofmzUqnZlgRynuTZ9Qcg9U9HpMWvHNvmONK7JFBv49Yt03eKEI9uNMS9tJF5qsw5QqZ-RZSu0rCsqgw5pFEOu1KTd87X",
                title: "The Art of Negative Space in Modern Digital Interfaces",
                blogTag: "Design",
                articleSnippet:
                    "Understanding how silence and void in design can amplify the impact of content and provide...",
                blogDate: "Oct 24, 2024",
              ),
              FeaturedBlogCard(
                imageUrl:
                    "https://lh3.googleusercontent.com/aida-public/AB6AXuBe3tsK9XNrc9wdeliJyJpElC_n1PAFG8pogrbHauvpRSWDnto0vgremLvalBmlTklkNgPFV7JfxE7k5N4YOPWydA1G-pBXPJHceWCrnefLpqTfuk8kng3impflpjpk_H5tE14J2NbnASiZL54C9Xpkcs2dXSL846zgL3jIjhi_ED76hHP7Vj2_SB6-FHfAdY-usGHiGyMchhMYQ0crPqtLv9l3oYPnIpvXVw65-oK0gL6Nd1Y2TZz0qmmt5YUkzDXCKHTq4Z1b",
                title: "Hardware for the Focused Mind",
                blogTag: "Tech",
                articleSnippet:
                    "Evaluating the tools that prioritize utility over distraction in an age of constant connectivity",
                blogDate: "Oct 22, 2024",
              ),
            ],
          ),
        ),
      ),
    );
  }
}
