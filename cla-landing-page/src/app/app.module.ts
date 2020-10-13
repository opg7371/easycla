// Copyright The Linux Foundation and each contributor to CommunityBridge.
// SPDX-License-Identifier: MIT

import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { ClaConsoleSectionComponent } from './components/cla-console-section/cla-console-section.component';
import { ClaFooterComponent } from './components/cla-footer/cla-footer.component';
import { LfxHeaderComponent } from './components/lfx-header/lfx-header.component';
import { AuthService } from './core/services/auth.service';
import { HomeComponent } from './components/home/home.component';
import { PageNotFoundComponent } from './components/page-not-found/page-not-found.component';
import { StorageService } from './core/services/storage.service';

@NgModule({
  declarations: [
    AppComponent,
    ClaConsoleSectionComponent,
    ClaFooterComponent,
    LfxHeaderComponent,
    HomeComponent,
    PageNotFoundComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule
  ],
  providers: [AuthService, StorageService],
  bootstrap: [AppComponent]
})
export class AppModule { }
