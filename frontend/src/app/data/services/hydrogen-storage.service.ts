import { Injectable, inject } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { HydrogenStorage } from '../interfaces/hydrogen-storage.interface';

@Injectable({
  providedIn: 'root'
})
export class HydrogenStorageService {
  http = inject(HttpClient)
  private apiUrl = 'http://192.168.1.199/api/';  

  getData() {
    return this.http.get<HydrogenStorage[]>(`${this.apiUrl}v1/hydrogen_storage`)
  }
}
