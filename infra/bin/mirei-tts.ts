#!/usr/bin/env node
import "source-map-support/register";
import * as cdk from "@aws-cdk/core";
import { MireiTTSStack } from "../lib/mirei-tts-stack";

const app = new cdk.App();
new MireiTTSStack(app, "MireiTTSStack", {
  env: {
    account: process.env.CDK_DEFAULT_ACCOUNT,
    region: process.env.CDK_DEFAULT_REGION,
  },
});
