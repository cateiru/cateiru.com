# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## 概要

cateiru.com の個人サイト。Next.js 16 + React 19 で構築し、本番は vinext 経由で Cloudflare Workers にデプロイする静的サイト。

## コマンド

```bash
# 開発サーバー起動（Next.js）
pnpm dev

# 開発サーバー起動（vinext）
pnpm dev:vinext

# 本番ビルド（vinext）
pnpm build:vinext

# vinext の本番ビルドをローカルで確認
pnpm start:vinext

# Cloudflare Workers へデプロイ（vinext）
pnpm deploy:vinext

# ローカルでのビルドプレビュー（OpenNext、legacy）
pnpm preview

# Cloudflare Workers へデプロイ（OpenNext、legacy）
pnpm deploy

# Lint / フォーマットチェック
pnpm check

# Lint / フォーマット自動修正
pnpm check:fix

# Cloudflare 環境型定義の生成
pnpm cf-typegen
```

## アーキテクチャ

- **フレームワーク**: Next.js 16 (App Router)
- **デプロイ先**: Cloudflare Workers（vinext 経由。OpenNext は legacy として併存）
- **Lint / フォーマット**: Biome（ESLint・Prettier の代替）
- **CSS**: CSS Modules + PostCSS (postcss-preset-env)
- **フォント**: `@fontsource/line-seed-jp`

### デプロイフロー

本番は Cloudflare Workers Builds でビルド・デプロイする。ビルドコマンドは `pnpm run build:vinext`、デプロイコマンドは `pnpm exec wrangler deploy --config dist/server/wrangler.json`、バージョンコマンドは `pnpm exec wrangler versions upload --config dist/server/wrangler.json`。`vite build` が `dist/server/wrangler.json` を生成し、実際のデプロイにはそちらを使う（ルートの `wrangler.jsonc` は vite / vinext の入力用）。

OpenNext の経路は legacy として残している。`pnpm preview` / `pnpm deploy` は専用の `wrangler.opennext.jsonc`（エントリは `.open-next/worker.js`）を使う。

### コードスタイル

Biome の設定（`biome.json`）に従う:

- インデント: スペース 2 文字
- クォート: ダブルクォート
- セミコロン: あり
- trailing comma: ES5 準拠

### パスエイリアス

`@/` は `components/` や `app/` などプロジェクトルートへのエイリアスとして使用可能（`tsconfig.json` による）。
