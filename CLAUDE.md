# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## 概要

cateiru.com の個人サイト。Next.js 15 + React 19 で構築し、Cloudflare Workers にデプロイする静的サイト。

## コマンド

```bash
# 開発サーバー起動
pnpm dev

# ローカルでのビルドプレビュー（Cloudflare 向け）
pnpm preview

# Cloudflare Workers へデプロイ
pnpm deploy

# Lint / フォーマットチェック
pnpm check

# Lint / フォーマット自動修正
pnpm check:fix

# Cloudflare 環境型定義の生成
pnpm cf-typegen
```

## アーキテクチャ

- **フレームワーク**: Next.js 15 (App Router)
- **デプロイ先**: Cloudflare Workers（`@opennextjs/cloudflare` 経由）
- **Lint / フォーマット**: Biome（ESLint・Prettier の代替）
- **CSS**: CSS Modules + PostCSS (postcss-preset-env)
- **フォント**: `@fontsource/line-seed-jp`

### デプロイフロー

`opennextjs-cloudflare` が Next.js のビルド成果物を Cloudflare Workers 向けに変換する。ビルド後の成果物は `.open-next/` に出力される。`wrangler.jsonc` でワーカー名・静的アセットディレクトリを設定している。

### コードスタイル

Biome の設定（`biome.json`）に従う:

- インデント: スペース 2 文字
- クォート: ダブルクォート
- セミコロン: あり
- trailing comma: ES5 準拠

### パスエイリアス

`@/` は `components/` や `app/` などプロジェクトルートへのエイリアスとして使用可能（`tsconfig.json` による）。
