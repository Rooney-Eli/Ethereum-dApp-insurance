package com.example.off_chain_oracle_client.ui

import androidx.hilt.lifecycle.ViewModelInject
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.example.off_chain_oracle_client.model.Validity
import com.example.off_chain_oracle_client.repository.MainRepository
import kotlinx.coroutines.launch


class MainViewModel
@ViewModelInject
constructor(
    private val mainRepo: MainRepository
) : ViewModel() {

    fun setValidityStateEvent(validityStateEvent: ValidityStateEvent<Validity>) {
        viewModelScope.launch {
            when (validityStateEvent) {
                is ValidityStateEvent.ValidityEvent<Validity> -> {
                    mainRepo.setValidity(validityStateEvent.data)
                }
                is ValidityStateEvent.None -> {
                    //Nothing
                }
            }
        }
    }
}


sealed class ValidityStateEvent<out R> {
    data class ValidityEvent<out T>(val data: T): ValidityStateEvent<T>()
    object None: ValidityStateEvent<Nothing>()
}