#!/usr/bin/env node
import "source-map-support/register";
import * as cdk from "@aws-cdk/core";
import { MireiTTSStack } from "../lib/mirei-tts-stack";

const app = new cdk.App();
new MireiTTSStack(app, "MireiTTSStack");
